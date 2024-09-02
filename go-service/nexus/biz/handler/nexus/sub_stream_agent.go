// Package nexus @Author Adrian.Wang 2024/9/1 17:35:00
// 是次级 agent 的流监控
package nexus

import (
	"fmt"
	"github.com/AdrianWangs/ai-nexus/go-service/nexus/kitex_gen/nexus_microservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/kr/pretty"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/packages/ssestream"
)

// ForwardResponseForSubNexus 跟踪次级 Nexus 的流
func (sa *StreamAgent) ForwardResponseForSubNexus(source *ssestream.Stream[openai.ChatCompletionChunk], target nexus_microservice.NexusService_AskServerServer, mainStreamAgent *StreamAgent) {

	// 开始对话,使用代理模式进行对话
	for source.Next() {

		event := source.Current()
		// 如果本轮对话没有任何回复就不需要进行其他额外的操作了
		if len(event.Choices) <= 0 {
			klog.Info("好像没对话内容...")
			pretty.Println(event)
			continue
		}

		// 将 openai 传过来的数据转化成我们网站对应的 response 格式
		askResponse := Event2response(event)

		// 不输出函数相关的内容，等函数生成完毕，才开始调用
		if len(askResponse.Choices[0].Message[0].ToolCalls) > 0 {
			continue
		}

		// 监控流，在监控过程中函数生成成功的那一刻进行函数调用
		sa.MonitorForSubNexus(event, target, mainStreamAgent)

		// 监控完以后不出意外就该转发刚刚的对话了
		err := target.Send(askResponse)
		if err != nil {
			fmt.Println("次级 ai：ForwardResponseForSubNexus--> 发送流给用户: ", err)
		}
	}

	if err := source.Err(); err != nil {
		klog.Error("次级 ai：ForwardResponseForSubNexus error:", err)
		sa.isStop = true
	}
}

// MonitorForSubNexus 监控流的请求,并执行相关函数调用
func (sa *StreamAgent) MonitorForSubNexus(event openai.ChatCompletionChunk, target nexus_microservice.NexusService_AskServerServer, mainStreamAgent *StreamAgent) {

	// 结束对话
	if event.Choices[0].FinishReason == openai.ChatCompletionChunkChoicesFinishReasonStop {
		// 结束本轮对话
		sa.EndConversation()
		return
	}

	// 当函数调用相关的参数生成完毕后，进行函数调用
	if event.Choices[0].FinishReason == openai.ChatCompletionChunkChoicesFinishReasonFunctionCall ||
		event.Choices[0].FinishReason == openai.ChatCompletionChunkChoicesFinishReasonToolCalls {

		finishReason := string(event.Choices[0].FinishReason)
		functionCallResponse := sa.GenerateToolMessageResponse(finishReason)
		// 监控完以后该转发刚刚的对话了
		err := target.Send(functionCallResponse)
		if err != nil {
			fmt.Println("MonitorForSubNexus--> 发送给用户的响应 :    执行错误: ", err)
		}

		// 调用函数
		sa.CallFunctionForSubNexus(target, mainStreamAgent)
		return
	}

	delta := event.Choices[0].Delta
	if delta.Content != "" {
		// 打印对话内容
		fmt.Print(delta.Content)
		sa.content += delta.Content
	}

	// 没有调用,直接返回
	if len(delta.ToolCalls) <= 0 {
		return
	}

	toolCall := delta.ToolCalls[0]
	// 判断是否是函数调用
	if toolCall.Type != openai.ChatCompletionChunkChoicesDeltaToolCallsTypeFunction {
		return
	}

	// 完善函数调用相关的信息，也就是切片组合成完整信息
	sa.CompleteFunctionCall(toolCall)
}

// CallFunctionForSubNexus 调用函数
func (sa *StreamAgent) CallFunctionForSubNexus(target nexus_microservice.NexusService_AskServerServer, mainStreamAgent *StreamAgent) {

	// 执行函数
	res, err := sa.DoFunctionForSubNexus(target)
	if err != nil {
		klog.Error("函数调用失败:", err)
		// 清空上下文，防止前面流影响后面的操作
		sa.ClearContext()
		return
	}

	// 这里应该是固定的 openai 格式（目前）
	if sa._type == "" {
		sa._type = "tool"
	}

	// 返回工具调用结果作为工具调用消息，插入到消息队列中
	toolMessage := sa.GenerateToolMessage(res)
	// 返回机器人的消息，插入到消息队列中
	assistantMessages := sa.GenerateAssistantMessage()
	// 将消息添加到消息列表中
	sa.messages = append(sa.messages, assistantMessages, toolMessage)

	// 添加消息到主 ai 的消息队列中
	mainStreamAgent.AddMessage(assistantMessages)
	mainStreamAgent.AddMessage(toolMessage)

	// 清空上下文，防止前面流影响后面的操作
	sa.ClearContext()

}

// DoFunctionForSubNexus 执行函数
func (sa *StreamAgent) DoFunctionForSubNexus(target nexus_microservice.NexusService_AskServerServer) (string, error) {

	fmt.Println("==========")
	fmt.Println("调用函数:", sa.functionName)
	fmt.Println("调用参数:", sa.functionArguments)
	fmt.Println("调用结果:", "")
	fmt.Println("==========")

	// 将方法转化给次级 ai 进行调用
	return "", nil
}
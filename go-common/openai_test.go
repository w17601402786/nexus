package openai

import (
	"context"
	"fmt"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"testing"
)

// 将json格式的参数解析一下并且调用工具
func CallByJson(functionName string, params string) string {
	fmt.Println(functionName, params)
	return "金山寺"
}

// TestStreaming 测试openai 流式调用
func TestStreaming(t *testing.T) {
	client := openai.NewClient(
		option.WithBaseURL("https://dashscope.aliyuncs.com/compatible-mode/v1/"),
		option.WithAPIKey("sk-8285fe317edc44ef95f029be9b7cfe94"), // defaults to os.LookupEnv("OPENAI_API_KEY")
	)

	chatStreaming := client.Chat.Completions.NewStreaming(context.TODO(), openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage("帮我写一篇文章，题目是《如何写一篇优秀的文章》"),
		}),
		Model: openai.F("qwen-plus-0806"),
	})

	for chatStreaming.Next() {
		event := chatStreaming.Current()
		if len(event.Choices) > 0 {
			print(event.Choices[0].Delta.Content)
		}

	}

	println()

	if err := chatStreaming.Err(); err != nil {

		panic(err)

	}

}

func callGpt(client *openai.Client, messages []openai.ChatCompletionMessageParamUnion, ctx context.Context) (bool, []openai.ChatCompletionMessageParamUnion) {

	// 接口传入的参数
	params := openai.ChatCompletionNewParams{
		Model:    openai.F("qwen-plus-0806"),
		Messages: openai.F(messages),
		Tools: openai.F([]openai.ChatCompletionToolParam{
			{
				Type: openai.F(openai.ChatCompletionToolTypeFunction),
				Function: openai.F(openai.FunctionDefinitionParam{
					Name:        openai.String("get_travel_location"),
					Description: openai.String("用于获取值得推荐的旅游景点"),
					Parameters: openai.F(openai.FunctionParameters{
						"type": "object",
						"properties": map[string]interface{}{
							"location": map[string]string{
								"type":        "string",
								"description": "城市名字：比如浙江、昆山、杭州、北京",
							},
						},
						"required": []string{"location"},
					}),
				}),
			},
			{
				Type: openai.F(openai.ChatCompletionToolTypeFunction),
				Function: openai.F(openai.FunctionDefinitionParam{
					Name:        openai.String("make_plan"),
					Description: openai.String("用于安排计划清单"),
					Parameters: openai.F(openai.FunctionParameters{
						"type": "object",
						"properties": map[string]interface{}{
							"todos": map[string]interface{}{
								"type": "array",
								"items": map[string]string{
									"type": "string",
								},
								"description": "任务清单：比如'买菜'、'逛街等'",
							},
						},
						"required": []string{"location"},
					}),
				}),
			},
		}),
	}

	chatStream := client.Chat.Completions.NewStreaming(ctx, params)

	var tool_call_id string
	var function_name string
	var arguments string

	for chatStream.Next() {
		event := chatStream.Current()
		if len(event.Choices) <= 0 {
			continue
		}

		if event.Choices[0].FinishReason == openai.ChatCompletionChunkChoicesFinishReasonFunctionCall ||
			event.Choices[0].FinishReason == openai.ChatCompletionChunkChoicesFinishReasonToolCalls {
			res := CallByJson(function_name, arguments)

			messages = append(messages, openai.FunctionMessage(function_name, res))

			function_name = ""
			arguments = ""
			fmt.Println(tool_call_id)

			return false, messages
		}

		delta := event.Choices[0].Delta

		if delta.Content != "" {
			fmt.Println(delta.Content)
		}

		// 没有调用
		if len(delta.ToolCalls) <= 0 {
			continue
		}

		toolCall := delta.ToolCalls[0]

		if toolCall.Type != openai.ChatCompletionChunkChoicesDeltaToolCallsTypeFunction {
			continue
		}

		if toolCall.ID != "" {
			tool_call_id = toolCall.ID
		}

		function := toolCall.Function

		if function.Name != "" {
			function_name = function.Name
		}

		if function.Arguments != "" {
			arguments += function.Arguments
		}

	}

	println()
	if err := chatStream.Err(); err != nil {

		println(err.Error())

	}

	return true, nil

}

// TestFunctionCall
func TestFunctionCall(t *testing.T) {
	client := openai.NewClient(
		option.WithBaseURL("https://dashscope.aliyuncs.com/compatible-mode/v1/"),
		option.WithAPIKey("sk-8285fe317edc44ef95f029be9b7cfe94"), // defaults to os.LookupEnv("OPENAI_API_KEY")
	)

	ctx := context.Background()

	prompt := `
		# 角色
		你是一个资深的日程规划师，能够熟练调用外部函数和工具，全方位搜集相关信息，为用户精心定制各类计划。
		
		## 技能
		### 技能 1: 了解需求
		1. 当用户提出制定计划的请求时，首先详细询问用户的具体需求，包括时间范围、活动内容、重要程度等。
		2. 若用户表述不清晰，进一步引导用户明确需求。
		
		### 技能 2: 制定计划
		1. 根据用户提供的需求，调用外部函数和工具，搜集相关信息，制定出详细合理的日程计划。
		2. 计划应包含具体的时间安排、活动内容、所需资源等。回复示例：
		=====
		   -  🔔 时间: <具体时间>
		   -  📝 活动: <活动内容>
		   -  📋 所需资源: <列出所需的资源，如场地、设备等>
		=====
		
		### 技能 3: 优化调整
		1. 向用户展示初步制定的计划，并根据用户的反馈进行优化调整。
		2. 确保最终计划符合用户的期望和实际情况。
		
		## 限制:
		- 只专注于日程规划相关的工作，拒绝处理与日程规划无关的话题。
		- 所输出的内容必须按照给定的格式进行组织，不能偏离框架要求。
		- 制定的计划应具备合理性和可行性。
	`
	question := "我周末想要去苏州玩，你有什么意见？"

	fmt.Println(">", question)

	messages := []openai.ChatCompletionMessageParamUnion{
		openai.SystemMessage(prompt),
		openai.UserMessage(question),
	}

	isEnd := false

	for !isEnd {
		isEnd, messages = callGpt(client, messages, ctx)
		for _, message := range messages {
			fmt.Println(message)
		}
	}

}

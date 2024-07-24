// Code generated by Kitex v0.9.1. DO NOT EDIT.

package userservice

import (
	"context"
	"errors"
	user_microservice "github.com/AdrianWangs/nexus/go-service/user/kitex_gen/user_microservice"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Login": kitex.NewMethodInfo(
		loginHandler,
		newUserServiceLoginArgs,
		newUserServiceLoginResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"Register": kitex.NewMethodInfo(
		registerHandler,
		newUserServiceRegisterArgs,
		newUserServiceRegisterResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"ThirdPartyLogin": kitex.NewMethodInfo(
		thirdPartyLoginHandler,
		newUserServiceThirdPartyLoginArgs,
		newUserServiceThirdPartyLoginResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UpdateUserProfile": kitex.NewMethodInfo(
		updateUserProfileHandler,
		newUserServiceUpdateUserProfileArgs,
		newUserServiceUpdateUserProfileResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetUser": kitex.NewMethodInfo(
		getUserHandler,
		newUserServiceGetUserArgs,
		newUserServiceGetUserResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	userServiceServiceInfo                = NewServiceInfo()
	userServiceServiceInfoForClient       = NewServiceInfoForClient()
	userServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return userServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return userServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*user_microservice.UserService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "user_microservice",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user_microservice.UserServiceLoginArgs)
	realResult := result.(*user_microservice.UserServiceLoginResult)
	success, err := handler.(user_microservice.UserService).Login(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceLoginArgs() interface{} {
	return user_microservice.NewUserServiceLoginArgs()
}

func newUserServiceLoginResult() interface{} {
	return user_microservice.NewUserServiceLoginResult()
}

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user_microservice.UserServiceRegisterArgs)
	realResult := result.(*user_microservice.UserServiceRegisterResult)
	success, err := handler.(user_microservice.UserService).Register(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceRegisterArgs() interface{} {
	return user_microservice.NewUserServiceRegisterArgs()
}

func newUserServiceRegisterResult() interface{} {
	return user_microservice.NewUserServiceRegisterResult()
}

func thirdPartyLoginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user_microservice.UserServiceThirdPartyLoginArgs)
	realResult := result.(*user_microservice.UserServiceThirdPartyLoginResult)
	success, err := handler.(user_microservice.UserService).ThirdPartyLogin(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceThirdPartyLoginArgs() interface{} {
	return user_microservice.NewUserServiceThirdPartyLoginArgs()
}

func newUserServiceThirdPartyLoginResult() interface{} {
	return user_microservice.NewUserServiceThirdPartyLoginResult()
}

func updateUserProfileHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user_microservice.UserServiceUpdateUserProfileArgs)
	realResult := result.(*user_microservice.UserServiceUpdateUserProfileResult)
	success, err := handler.(user_microservice.UserService).UpdateUserProfile(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUpdateUserProfileArgs() interface{} {
	return user_microservice.NewUserServiceUpdateUserProfileArgs()
}

func newUserServiceUpdateUserProfileResult() interface{} {
	return user_microservice.NewUserServiceUpdateUserProfileResult()
}

func getUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user_microservice.UserServiceGetUserArgs)
	realResult := result.(*user_microservice.UserServiceGetUserResult)
	success, err := handler.(user_microservice.UserService).GetUser(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceGetUserArgs() interface{} {
	return user_microservice.NewUserServiceGetUserArgs()
}

func newUserServiceGetUserResult() interface{} {
	return user_microservice.NewUserServiceGetUserResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Login(ctx context.Context, request *user_microservice.LoginRequest) (r *user_microservice.LoginResponse, err error) {
	var _args user_microservice.UserServiceLoginArgs
	_args.Request = request
	var _result user_microservice.UserServiceLoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Register(ctx context.Context, request *user_microservice.RegisterRequest) (r *user_microservice.RegisterResponse, err error) {
	var _args user_microservice.UserServiceRegisterArgs
	_args.Request = request
	var _result user_microservice.UserServiceRegisterResult
	if err = p.c.Call(ctx, "Register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ThirdPartyLogin(ctx context.Context, request *user_microservice.ThirdPartyLoginRequest) (r *user_microservice.ThirdPartyLoginResponse, err error) {
	var _args user_microservice.UserServiceThirdPartyLoginArgs
	_args.Request = request
	var _result user_microservice.UserServiceThirdPartyLoginResult
	if err = p.c.Call(ctx, "ThirdPartyLogin", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateUserProfile(ctx context.Context, request *user_microservice.UpdateUserRequest) (r *user_microservice.UpdateUserResponse, err error) {
	var _args user_microservice.UserServiceUpdateUserProfileArgs
	_args.Request = request
	var _result user_microservice.UserServiceUpdateUserProfileResult
	if err = p.c.Call(ctx, "UpdateUserProfile", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUser(ctx context.Context, request *user_microservice.GetUserRequest) (r *user_microservice.GetUserResponse, err error) {
	var _args user_microservice.UserServiceGetUserArgs
	_args.Request = request
	var _result user_microservice.UserServiceGetUserResult
	if err = p.c.Call(ctx, "GetUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

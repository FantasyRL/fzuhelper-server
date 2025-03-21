/*
Copyright 2024 The west2-online Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kitex v0.12.1. DO NOT EDIT.

package userservice

import (
	"context"
	"errors"

	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"

	user "github.com/west2-online/fzuhelper-server/kitex_gen/user"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"GetLoginData": kitex.NewMethodInfo(
		getLoginDataHandler,
		newUserServiceGetLoginDataArgs,
		newUserServiceGetLoginDataResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetUserInfo": kitex.NewMethodInfo(
		getUserInfoHandler,
		newUserServiceGetUserInfoArgs,
		newUserServiceGetUserInfoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetGetLoginDataForYJSY": kitex.NewMethodInfo(
		getGetLoginDataForYJSYHandler,
		newUserServiceGetGetLoginDataForYJSYArgs,
		newUserServiceGetGetLoginDataForYJSYResult,
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

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return userServiceServiceInfoForStreamClient
}

// for client
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
	handlerType := (*user.UserService)(nil)
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
		"PackageName": "user",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.12.1",
		Extra:           extra,
	}
	return svcInfo
}

func getLoginDataHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceGetLoginDataArgs)
	realResult := result.(*user.UserServiceGetLoginDataResult)
	success, err := handler.(user.UserService).GetLoginData(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceGetLoginDataArgs() interface{} {
	return user.NewUserServiceGetLoginDataArgs()
}

func newUserServiceGetLoginDataResult() interface{} {
	return user.NewUserServiceGetLoginDataResult()
}

func getUserInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceGetUserInfoArgs)
	realResult := result.(*user.UserServiceGetUserInfoResult)
	success, err := handler.(user.UserService).GetUserInfo(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceGetUserInfoArgs() interface{} {
	return user.NewUserServiceGetUserInfoArgs()
}

func newUserServiceGetUserInfoResult() interface{} {
	return user.NewUserServiceGetUserInfoResult()
}

func getGetLoginDataForYJSYHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceGetGetLoginDataForYJSYArgs)
	realResult := result.(*user.UserServiceGetGetLoginDataForYJSYResult)
	success, err := handler.(user.UserService).GetGetLoginDataForYJSY(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceGetGetLoginDataForYJSYArgs() interface{} {
	return user.NewUserServiceGetGetLoginDataForYJSYArgs()
}

func newUserServiceGetGetLoginDataForYJSYResult() interface{} {
	return user.NewUserServiceGetGetLoginDataForYJSYResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetLoginData(ctx context.Context, req *user.GetLoginDataRequest) (r *user.GetLoginDataResponse, err error) {
	var _args user.UserServiceGetLoginDataArgs
	_args.Req = req
	var _result user.UserServiceGetLoginDataResult
	if err = p.c.Call(ctx, "GetLoginData", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserInfo(ctx context.Context, request *user.GetUserInfoRequest) (r *user.GetUserInfoResponse, err error) {
	var _args user.UserServiceGetUserInfoArgs
	_args.Request = request
	var _result user.UserServiceGetUserInfoResult
	if err = p.c.Call(ctx, "GetUserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetGetLoginDataForYJSY(ctx context.Context, request *user.GetLoginDataForYJSYRequest) (r *user.GetLoginDataForYJSYResponse, err error) {
	var _args user.UserServiceGetGetLoginDataForYJSYArgs
	_args.Request = request
	var _result user.UserServiceGetGetLoginDataForYJSYResult
	if err = p.c.Call(ctx, "GetGetLoginDataForYJSY", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

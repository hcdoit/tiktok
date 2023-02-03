// Code generated by Kitex v0.4.4. DO NOT EDIT.

package socialservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	social "github.com/hcdoit/tiktok/kitex_gen/social"
)

func serviceInfo() *kitex.ServiceInfo {
	return socialServiceServiceInfo
}

var socialServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "SocialService"
	handlerType := (*social.SocialService)(nil)
	methods := map[string]kitex.MethodInfo{
		"RelationAction":  kitex.NewMethodInfo(relationActionHandler, newSocialServiceRelationActionArgs, newSocialServiceRelationActionResult, false),
		"GetFollowList":   kitex.NewMethodInfo(getFollowListHandler, newSocialServiceGetFollowListArgs, newSocialServiceGetFollowListResult, false),
		"GetFollowerList": kitex.NewMethodInfo(getFollowerListHandler, newSocialServiceGetFollowerListArgs, newSocialServiceGetFollowerListResult, false),
		"GetFriendList":   kitex.NewMethodInfo(getFriendListHandler, newSocialServiceGetFriendListArgs, newSocialServiceGetFriendListResult, false),
		"GetRelationInfo": kitex.NewMethodInfo(getRelationInfoHandler, newSocialServiceGetRelationInfoArgs, newSocialServiceGetRelationInfoResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "social",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func relationActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*social.SocialServiceRelationActionArgs)
	realResult := result.(*social.SocialServiceRelationActionResult)
	success, err := handler.(social.SocialService).RelationAction(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSocialServiceRelationActionArgs() interface{} {
	return social.NewSocialServiceRelationActionArgs()
}

func newSocialServiceRelationActionResult() interface{} {
	return social.NewSocialServiceRelationActionResult()
}

func getFollowListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*social.SocialServiceGetFollowListArgs)
	realResult := result.(*social.SocialServiceGetFollowListResult)
	success, err := handler.(social.SocialService).GetFollowList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSocialServiceGetFollowListArgs() interface{} {
	return social.NewSocialServiceGetFollowListArgs()
}

func newSocialServiceGetFollowListResult() interface{} {
	return social.NewSocialServiceGetFollowListResult()
}

func getFollowerListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*social.SocialServiceGetFollowerListArgs)
	realResult := result.(*social.SocialServiceGetFollowerListResult)
	success, err := handler.(social.SocialService).GetFollowerList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSocialServiceGetFollowerListArgs() interface{} {
	return social.NewSocialServiceGetFollowerListArgs()
}

func newSocialServiceGetFollowerListResult() interface{} {
	return social.NewSocialServiceGetFollowerListResult()
}

func getFriendListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*social.SocialServiceGetFriendListArgs)
	realResult := result.(*social.SocialServiceGetFriendListResult)
	success, err := handler.(social.SocialService).GetFriendList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSocialServiceGetFriendListArgs() interface{} {
	return social.NewSocialServiceGetFriendListArgs()
}

func newSocialServiceGetFriendListResult() interface{} {
	return social.NewSocialServiceGetFriendListResult()
}

func getRelationInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*social.SocialServiceGetRelationInfoArgs)
	realResult := result.(*social.SocialServiceGetRelationInfoResult)
	success, err := handler.(social.SocialService).GetRelationInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSocialServiceGetRelationInfoArgs() interface{} {
	return social.NewSocialServiceGetRelationInfoArgs()
}

func newSocialServiceGetRelationInfoResult() interface{} {
	return social.NewSocialServiceGetRelationInfoResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) RelationAction(ctx context.Context, req *social.RelationActionRequest) (r *social.RelationActionResponse, err error) {
	var _args social.SocialServiceRelationActionArgs
	_args.Req = req
	var _result social.SocialServiceRelationActionResult
	if err = p.c.Call(ctx, "RelationAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFollowList(ctx context.Context, req *social.RelationListRequest) (r *social.RelationListResponse, err error) {
	var _args social.SocialServiceGetFollowListArgs
	_args.Req = req
	var _result social.SocialServiceGetFollowListResult
	if err = p.c.Call(ctx, "GetFollowList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFollowerList(ctx context.Context, req *social.RelationListRequest) (r *social.RelationListResponse, err error) {
	var _args social.SocialServiceGetFollowerListArgs
	_args.Req = req
	var _result social.SocialServiceGetFollowerListResult
	if err = p.c.Call(ctx, "GetFollowerList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFriendList(ctx context.Context, req *social.RelationListRequest) (r *social.RelationListResponse, err error) {
	var _args social.SocialServiceGetFriendListArgs
	_args.Req = req
	var _result social.SocialServiceGetFriendListResult
	if err = p.c.Call(ctx, "GetFriendList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetRelationInfo(ctx context.Context, req *social.RelationInfoRequest) (r *social.RelationInfoResponse, err error) {
	var _args social.SocialServiceGetRelationInfoArgs
	_args.Req = req
	var _result social.SocialServiceGetRelationInfoResult
	if err = p.c.Call(ctx, "GetRelationInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

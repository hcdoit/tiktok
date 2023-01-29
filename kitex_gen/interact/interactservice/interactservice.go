// Code generated by Kitex v0.4.4. DO NOT EDIT.

package interactservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	interact "github.com/hcdoit/tiktok/kitex_gen/interact"
)

func serviceInfo() *kitex.ServiceInfo {
	return interactServiceServiceInfo
}

var interactServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "InteractService"
	handlerType := (*interact.InteractService)(nil)
	methods := map[string]kitex.MethodInfo{
		"FavoriteAction":  kitex.NewMethodInfo(favoriteActionHandler, newInteractServiceFavoriteActionArgs, newInteractServiceFavoriteActionResult, false),
		"GetFavoriteList": kitex.NewMethodInfo(getFavoriteListHandler, newInteractServiceGetFavoriteListArgs, newInteractServiceGetFavoriteListResult, false),
		"CommentAction":   kitex.NewMethodInfo(commentActionHandler, newInteractServiceCommentActionArgs, newInteractServiceCommentActionResult, false),
		"GetCommentList":  kitex.NewMethodInfo(getCommentListHandler, newInteractServiceGetCommentListArgs, newInteractServiceGetCommentListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "interact",
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

func favoriteActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*interact.InteractServiceFavoriteActionArgs)
	realResult := result.(*interact.InteractServiceFavoriteActionResult)
	success, err := handler.(interact.InteractService).FavoriteAction(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newInteractServiceFavoriteActionArgs() interface{} {
	return interact.NewInteractServiceFavoriteActionArgs()
}

func newInteractServiceFavoriteActionResult() interface{} {
	return interact.NewInteractServiceFavoriteActionResult()
}

func getFavoriteListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*interact.InteractServiceGetFavoriteListArgs)
	realResult := result.(*interact.InteractServiceGetFavoriteListResult)
	success, err := handler.(interact.InteractService).GetFavoriteList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newInteractServiceGetFavoriteListArgs() interface{} {
	return interact.NewInteractServiceGetFavoriteListArgs()
}

func newInteractServiceGetFavoriteListResult() interface{} {
	return interact.NewInteractServiceGetFavoriteListResult()
}

func commentActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*interact.InteractServiceCommentActionArgs)
	realResult := result.(*interact.InteractServiceCommentActionResult)
	success, err := handler.(interact.InteractService).CommentAction(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newInteractServiceCommentActionArgs() interface{} {
	return interact.NewInteractServiceCommentActionArgs()
}

func newInteractServiceCommentActionResult() interface{} {
	return interact.NewInteractServiceCommentActionResult()
}

func getCommentListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*interact.InteractServiceGetCommentListArgs)
	realResult := result.(*interact.InteractServiceGetCommentListResult)
	success, err := handler.(interact.InteractService).GetCommentList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newInteractServiceGetCommentListArgs() interface{} {
	return interact.NewInteractServiceGetCommentListArgs()
}

func newInteractServiceGetCommentListResult() interface{} {
	return interact.NewInteractServiceGetCommentListResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) FavoriteAction(ctx context.Context, req *interact.FavoriteActionRequest) (r *interact.FavoriteActionResponse, err error) {
	var _args interact.InteractServiceFavoriteActionArgs
	_args.Req = req
	var _result interact.InteractServiceFavoriteActionResult
	if err = p.c.Call(ctx, "FavoriteAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFavoriteList(ctx context.Context, req *interact.FavoriteListRequest) (r *interact.FavoriteListResponse, err error) {
	var _args interact.InteractServiceGetFavoriteListArgs
	_args.Req = req
	var _result interact.InteractServiceGetFavoriteListResult
	if err = p.c.Call(ctx, "GetFavoriteList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CommentAction(ctx context.Context, req *interact.CommentActionRequest) (r *interact.CommentActionResponse, err error) {
	var _args interact.InteractServiceCommentActionArgs
	_args.Req = req
	var _result interact.InteractServiceCommentActionResult
	if err = p.c.Call(ctx, "CommentAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetCommentList(ctx context.Context, req *interact.CommentListRequest) (r *interact.CommentListResponse, err error) {
	var _args interact.InteractServiceGetCommentListArgs
	_args.Req = req
	var _result interact.InteractServiceGetCommentListResult
	if err = p.c.Call(ctx, "GetCommentList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
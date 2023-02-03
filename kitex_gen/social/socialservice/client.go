// Code generated by Kitex v0.4.4. DO NOT EDIT.

package socialservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	social "github.com/hcdoit/tiktok/kitex_gen/social"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	RelationAction(ctx context.Context, req *social.RelationActionRequest, callOptions ...callopt.Option) (r *social.RelationActionResponse, err error)
	GetFollowList(ctx context.Context, req *social.RelationListRequest, callOptions ...callopt.Option) (r *social.RelationListResponse, err error)
	GetFollowerList(ctx context.Context, req *social.RelationListRequest, callOptions ...callopt.Option) (r *social.RelationListResponse, err error)
	GetFriendList(ctx context.Context, req *social.RelationListRequest, callOptions ...callopt.Option) (r *social.RelationListResponse, err error)
	GetRelationInfo(ctx context.Context, req *social.RelationInfoRequest, callOptions ...callopt.Option) (r *social.RelationInfoResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kSocialServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kSocialServiceClient struct {
	*kClient
}

func (p *kSocialServiceClient) RelationAction(ctx context.Context, req *social.RelationActionRequest, callOptions ...callopt.Option) (r *social.RelationActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RelationAction(ctx, req)
}

func (p *kSocialServiceClient) GetFollowList(ctx context.Context, req *social.RelationListRequest, callOptions ...callopt.Option) (r *social.RelationListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFollowList(ctx, req)
}

func (p *kSocialServiceClient) GetFollowerList(ctx context.Context, req *social.RelationListRequest, callOptions ...callopt.Option) (r *social.RelationListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFollowerList(ctx, req)
}

func (p *kSocialServiceClient) GetFriendList(ctx context.Context, req *social.RelationListRequest, callOptions ...callopt.Option) (r *social.RelationListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFriendList(ctx, req)
}

func (p *kSocialServiceClient) GetRelationInfo(ctx context.Context, req *social.RelationInfoRequest, callOptions ...callopt.Option) (r *social.RelationInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetRelationInfo(ctx, req)
}

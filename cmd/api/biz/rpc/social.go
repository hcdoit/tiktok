package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/hcdoit/tiktok/kitex_gen/social"
	"github.com/hcdoit/tiktok/kitex_gen/social/socialservice"
	"github.com/hcdoit/tiktok/pkg/consts"
	"github.com/hcdoit/tiktok/pkg/mw"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var socialClient socialservice.Client

func initSocial() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.ApiServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := socialservice.NewClient(
		consts.SocialServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	socialClient = c
}

func GetFriendList(ctx context.Context, req *social.RelationListRequest) (resp *social.RelationListResponse, err error) {
	resp, err = socialClient.GetFriendList(ctx, req)
	return resp, err
}
func GetFollowerList(ctx context.Context, req *social.RelationListRequest) (resp *social.RelationListResponse, err error) {
	resp, err = socialClient.GetFollowerList(ctx, req)
	return resp, err
}
func GetFollowList(ctx context.Context, req *social.RelationListRequest) (resp *social.RelationListResponse, err error) {
	resp, err = socialClient.GetFollowList(ctx, req)
	return resp, err
}

func RelationAction(ctx context.Context, req *social.RelationActionRequest) (resp *social.RelationActionResponse, err error) {
	resp, err = socialClient.RelationAction(ctx, req)
	return resp, err
}

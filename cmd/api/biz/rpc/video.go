package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/hcdoit/tiktok/kitex_gen/video"
	"github.com/hcdoit/tiktok/kitex_gen/video/videoservice"
	"github.com/hcdoit/tiktok/pkg/consts"
	"github.com/hcdoit/tiktok/pkg/mw"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var videoClient videoservice.Client

func initVideo() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.ApiServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := videoservice.NewClient(
		consts.VideoServiceName,
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
	videoClient = c
}

func PublishAction(ctx context.Context, req *video.PublishActionRequest) (*video.PublishActionResponse, error) {
	resp, err := videoClient.PublishAction(ctx, req)
	return resp, err
}

func GetFeed(ctx context.Context, req *video.FeedRequest) (*video.FeedResponse, error) {
	resp, err := videoClient.GetFeed(ctx, req)
	return resp, err
}

func GetPublishList(ctx context.Context, req *video.PublishListRequest) (*video.PublishListResponse, error) {
	resp, err := videoClient.GetPublishList(ctx, req)
	return resp, err
}

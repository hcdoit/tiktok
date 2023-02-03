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
		provider.WithServiceName(consts.InteractServiceName),
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
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.InteractServiceName}),
	)
	if err != nil {
		panic(err)
	}
	videoClient = c
}

func GetVideo(ctx context.Context, req *video.VideoRequest) (*video.VideoResponse, error) {
	resp, err := videoClient.GetVideo(ctx, req)
	return resp, err
}

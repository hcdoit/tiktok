package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/hcdoit/tiktok/kitex_gen/interact"
	"github.com/hcdoit/tiktok/kitex_gen/interact/interactservice"
	"github.com/hcdoit/tiktok/pkg/consts"
	"github.com/hcdoit/tiktok/pkg/errno"
	"github.com/hcdoit/tiktok/pkg/mw"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var interactClient interactservice.Client

func initInteract() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.VideoServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := interactservice.NewClient(
		consts.InteractServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.VideoServiceName}),
	)
	if err != nil {
		panic(err)
	}
	interactClient = c
}

func GetInteract(ctx context.Context, req *interact.VideoInteractRequest) (*interact.VideoInteractResponse, error) {
	resp, err := interactClient.GetVideoInteract(ctx, req)
	if err != nil || resp.StatusCode != errno.Success.ErrCode {
		klog.Error("video get interact error")
		return nil, errno.ResourceNotFound
	}

	return resp, nil
}

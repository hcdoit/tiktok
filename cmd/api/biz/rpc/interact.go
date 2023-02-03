package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/hcdoit/tiktok/kitex_gen/interact"
	"github.com/hcdoit/tiktok/kitex_gen/interact/interactservice"
	"github.com/hcdoit/tiktok/pkg/consts"
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
		provider.WithServiceName(consts.ApiServiceName),
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
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	interactClient = c
}

func CommentAction(ctx context.Context, req *interact.CommentActionRequest) (*interact.CommentActionResponse, error) {
	resp, err := interactClient.CommentAction(ctx, req)
	return resp, err
}
func GetCommentList(ctx context.Context, req *interact.CommentListRequest) (*interact.CommentListResponse, error) {
	resp, err := interactClient.GetCommentList(ctx, req)
	return resp, err
}

func FavoriteAction(ctx context.Context, req *interact.FavoriteActionRequest) (*interact.FavoriteActionResponse, error) {
	resp, err := interactClient.FavoriteAction(ctx, req)
	return resp, err
}
func GetFavoriteList(ctx context.Context, req *interact.FavoriteListRequest) (*interact.FavoriteListResponse, error) {
	resp, err := interactClient.GetFavoriteList(ctx, req)
	return resp, err
}

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
		provider.WithServiceName(consts.UserServiceName),
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
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.UserServiceName}),
	)
	if err != nil {
		panic(err)
	}
	socialClient = c
}

func GetRelationInfo(ctx context.Context, req *social.RelationInfoRequest) (resp *social.RelationInfoResponse, err error) {
	resp, err = socialClient.GetRelationInfo(ctx, req)
	return resp, err
}

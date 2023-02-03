package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/hcdoit/tiktok/cmd/video/dal"
	"github.com/hcdoit/tiktok/cmd/video/rpc"
	"github.com/hcdoit/tiktok/kitex_gen/video/videoservice"
	"github.com/hcdoit/tiktok/pkg/consts"
	"github.com/hcdoit/tiktok/pkg/mw"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"net"
)

func Init() {
	dal.Init()
	rpc.Init()
	klog.SetLogger(kitexlogrus.NewLogger())
	klog.SetLevel(consts.KLoggerLevel)
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr(consts.TCP, consts.VideoServiceAddr)
	if err != nil {
		panic(err)
	}
	Init()
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.VideoServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	svr := videoservice.NewServer(new(VideoServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithMuxTransport(),
		server.WithMiddleware(mw.CommonMiddleware),
		server.WithMiddleware(mw.ServerMiddleware),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.VideoServiceName}),
	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}

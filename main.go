package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/kpyramid/bitcoin-inscribe/proto"
	"github.com/kpyramid/bitcoin-inscribe/rpc"
	"github.com/kpyramid/bitcoin-inscribe/types"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"io"
	"os"
)

func init() {
	os.Mkdir("log", 0777)
	file, err := os.OpenFile(fmt.Sprintf("log/%s.log", "main"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("create log file failed, error: %s", err)
	}

	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)
	log.SetReportCaller(true)
}

func main() {
	svc := types.GetServiceContext()
	_ = svc

	go func() { types.Server(registerGrpc, registerGateway) }()

	if svc.Config.UseCoroutine {

	}

	// wait
	types.WaitQuitSignal(svc)
}

func registerGrpc(svc *grpc.Server) {
	// create grpc server
	proto.RegisterInscribeServiceServer(svc, rpc.Inscribe{})
}

func registerGateway(ctx context.Context, mux *runtime.ServeMux, endPoint string, option []grpc.DialOption) error {
	err := proto.RegisterInscribeServiceHandlerFromEndpoint(ctx, mux, endPoint, option)
	if err != nil {
		log.WithField("error", err).Panic("register grpc gateway server failed")
	}
	return nil
}

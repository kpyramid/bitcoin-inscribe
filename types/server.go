package types

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	uperror "github.com/kpyramid/bitcoin-inscribe/types/error"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type Validator interface {
	Validate() error
}

type GrpcGreeterRegisterFn func(*grpc.Server)
type GrpcGatewayRegisterFn func(ctx context.Context, mux *runtime.ServeMux, endPoint string, option []grpc.DialOption) error

func Server(grpcFn GrpcGreeterRegisterFn, gatewayFn GrpcGatewayRegisterFn, interceptorFunc ...grpc.UnaryServerInterceptor) {
	cfg := GetConfig()
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	grpcEndpoint := fmt.Sprintf("%s:%d", cfg.Address, cfg.GrpcPort)
	httpEndpoint := fmt.Sprintf("%s:%d", cfg.Address, cfg.HttpPort)

	recoveryOpt := grpc_recovery.WithRecoveryHandlerContext(
		func(ctx context.Context, p interface{}) error {
			log.Errorf("[PANIC] %s\n\n%s", p, string(debug.Stack()))
			return status.Errorf(codes.Unknown, "panic triggered: %v", p)
		},
	)
	opts := []grpc_recovery.Option{
		recoveryOpt,
	}

	reqValidatorInterceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if p, ok := req.(Validator); ok {
			if err := p.Validate(); err != nil {
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
		}
		return handler(ctx, req)
	}

	var interceptorArr []grpc.UnaryServerInterceptor
	interceptorArr = append(interceptorArr, reqValidatorInterceptor)
	if len(interceptorFunc) > 0 {
		interceptorArr = append(interceptorArr, interceptorFunc...)
	}
	var grpcSvc *grpc.Server

	grpcSvc = grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(opts...),
			grpc_middleware.ChainUnaryServer(interceptorArr...),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_recovery.StreamServerInterceptor(opts...),
		),
	)

	httpSvc := http.Server{}

	go func() {
		listen, err := net.Listen("tcp", grpcEndpoint)
		if err != nil {
			log.WithField("error", err).Fatal("listen failed")
		}

		// create grpc server
		grpcFn(grpcSvc)

		if err := grpcSvc.Serve(listen); err != nil {
			log.WithField("error", err).Fatal("serve grpc server failed")
		}
	}()

	go func() {
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		// http server
		listen, err := net.Listen("tcp", httpEndpoint)
		if err != nil {
			log.WithField("error", err).Fatal("start grpc gateway server failed")
		}

		// Register gRPC server endpoint
		mux := runtime.NewServeMux(
			runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}),
			runtime.WithProtoErrorHandler(protoErrorHandle),
		)
		opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithNoProxy()}
		if err := gatewayFn(ctx, mux, grpcEndpoint, opts); err != nil {
			log.WithField("error", err).Panic("register grpc gateway server failed")
		}

		httpSvc.Handler = mux
		httpSvc.Addr = httpEndpoint

		// Start http server
		if err := httpSvc.Serve(listen); err != nil {
			log.WithField("error", err).Panic("start grpc gateway server failed")
		}
	}()

	log.WithFields(log.Fields{"address": cfg.Address, "port": cfg.GrpcPort}).Info("grpc server listening")
	log.WithFields(log.Fields{"address": cfg.Address, "port": cfg.HttpPort}).Info("http server listening")

	<-stopChan
	grpcSvc.Stop()
	_ = httpSvc.Close()
}

func protoErrorHandle(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, writer http.ResponseWriter, request *http.Request, err error) {
	s, ok := status.FromError(err)
	if !ok {
		s = status.New(codes.Unknown, err.Error())
	}

	// set header
	writer.Header().Del("Trailer")
	writer.Header().Set("Context-Type", marshaler.ContentType())

	// set content
	body := &struct {
		ErrorCode    uperror.StatusErrorKey
		InternalCode codes.Code
		Message      string
		Details      []interface{}
	}{}

	authErr, err := uperror.AuthErrorFromString(s.Message())
	if err != nil {
		body.ErrorCode = uperror.StatusErrorUnknown
		body.InternalCode = s.Code()
		body.Message = s.Message()
		body.Details = s.Details()

	} else {
		body.ErrorCode = authErr.Code
		body.InternalCode = s.Code()
		body.Message = authErr.Message
		body.Details = s.Details()
	}

	buf, merr := marshaler.Marshal(body)
	if merr != nil {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = writer.Write([]byte(merr.Error()))
		return
	}
	if authErr != nil && authErr.Code == -20011 {
		writer.WriteHeader(http.StatusUnauthorized)
	} else {
		// set status
		writer.WriteHeader(runtime.HTTPStatusFromCode(codes.InvalidArgument))
	}

	_, _ = writer.Write(buf)
}

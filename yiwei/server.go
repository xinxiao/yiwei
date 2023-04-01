package main

import (
	"flag"
	"fmt"
	"net"
	"time"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcZap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcPrometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	grpcReflection "google.golang.org/grpc/reflection"
)

var (
	connTimeout = flag.Duration("connection_timeout", 10*time.Second, "grpc connection timeout")
	port        = flag.Int("port", 10725, "port to host yiwei server")
)

func CreateServer(log *zap.Logger) *grpc.Server {
	svr := grpc.NewServer(
		grpc.ConnectionTimeout(*connTimeout),
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(
				grpcZap.UnaryServerInterceptor(log),
				grpcPrometheus.UnaryServerInterceptor,
				grpcRecovery.UnaryServerInterceptor())),
		grpc.StreamInterceptor(
			grpcMiddleware.ChainStreamServer(
				grpcZap.StreamServerInterceptor(log),
				grpcPrometheus.StreamServerInterceptor,
				grpcRecovery.StreamServerInterceptor())))

	grpcReflection.Register(svr)
	grpcPrometheus.Register(svr)

	return svr
}

func Launch(log *zap.Logger, svr *grpc.Server, ec chan error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		ec <- err
		return
	}
	defer lis.Close()

	log.Info("launching grpc server", zap.Int("port", *port))
	defer svr.GracefulStop()
	if err := svr.Serve(lis); err != nil {
		ec <- err
	}
}

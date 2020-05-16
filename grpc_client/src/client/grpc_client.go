package client

import (
	"github.com/maei/shared_utils_go/logger"
	"google.golang.org/grpc"
	"log"
)

var GRPCClient grpcClientInterface = &grpcClient{}

type grpcClientInterface interface {
	SetClient() (*grpc.ClientConn, error)
}

type grpcClient struct{}

func (*grpcClient) SetClient() (*grpc.ClientConn, error) {
	logger.Info("starting gRPC Client")
	conn, err := grpc.Dial("0.0.0.0:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
		return nil, err
	}

	return conn, nil
}

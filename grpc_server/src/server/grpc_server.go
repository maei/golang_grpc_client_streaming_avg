package server

import (
	"github.com/maei/golang_grpc_client_streaming_avg/grpc_server/src/domain/greetpb"
	"github.com/maei/shared_utils_go/logger"
	"google.golang.org/grpc"
	"io"
	"net"
)

type server struct{}

var (
	s = grpc.NewServer()
)

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	logger.Info("LongGreet streaming request was invoked")
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			res := &greetpb.LongGreetResponse{
				Result: result,
			}
			return stream.SendAndClose(res)
		}
		if err != nil {
			logger.Error("Error while reading client stream", err)
		}
		firstName := req.GetGreeting().GetFirstName()
		result += "Hello " + firstName + "! "
	}
}

func StartGRPCServer() {
	logger.Info("gRPC greet-streaming started")

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		logger.Error("error while listening gRPC Server", err)
	}
	greetpb.RegisterGreetServiceServer(s, &server{})

	errServer := s.Serve(lis)
	if errServer != nil {
		logger.Error("error while serve gRPC Server", errServer)
	}
}

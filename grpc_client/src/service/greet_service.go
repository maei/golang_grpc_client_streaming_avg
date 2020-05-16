package service

import (
	"context"
	"fmt"
	"github.com/maei/golang_grpc_client_streaming_avg/grpc_client/src/client"
	"github.com/maei/golang_grpc_client_streaming_avg/grpc_client/src/domain/greetpb"
	"github.com/maei/shared_utils_go/logger"
	"time"
)

var GreetService greetServiceInterface = &greetService{}

type greetServiceInterface interface {
	LongGreet()
}

type greetService struct{}

func (*greetService) LongGreet() {
	conn, err := client.GRPCClient.SetClient()
	if err != nil {
		logger.Error("Error whilce creating GRPC Client", err)
	}
	defer conn.Close()
	c := greetpb.NewGreetServiceClient(conn)

	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{FirstName: "Matthias"},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{FirstName: "Sonia"},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{FirstName: "Heidi"},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{FirstName: "Maya"},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{FirstName: "Jochen"},
		},
	}

	stream, streamErr := c.LongGreet(context.Background())
	if streamErr != nil {
		logger.Error("Error while streaming data", err)
	}

	for _, req := range requests {
		logger.Info(fmt.Sprintf("sending %v\n", req))
		stream.Send(req)
		time.Sleep(2 * time.Second)
	}
	res, resErr := stream.CloseAndRecv()
	if resErr != nil {
		logger.Error("Error while waiting for grpc_server response", resErr)
	}
	fmt.Printf("GRPC-Server response %v\n", res)
}

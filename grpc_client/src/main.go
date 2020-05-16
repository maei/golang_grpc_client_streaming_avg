package main

import (
	"github.com/maei/golang_grpc_client_streaming_avg/grpc_client/src/service"
	"github.com/maei/shared_utils_go/logger"
)

func main() {
	logger.Info("Starting GRPC-Client")
	service.GreetService.LongGreet()
}

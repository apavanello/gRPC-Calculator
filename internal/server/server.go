package server

import (
	"log"
	"net"

	"github.com/apavanello/gRPC-Calculato/internal/Calculator"
	"github.com/apavanello/gRPC-Calculato/internal/MyVersion"
	"github.com/apavanello/gRPC-Calculato/pkg/calculatorpb"
	"google.golang.org/grpc"
)

func Run() {

	// bind gRPC port

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create gRPC server
	s := grpc.NewServer()

	// register gRPC server
	calculatorpb.RegisterCalculatorServiceServer(s, &Calculator.Server{})
	calculatorpb.RegisterVersionServiceServer(s, &MyVersion.Server{})

	// start gRPC server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/apavanello/gRPC-Calculato/pkg/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func gRPCCalculate(x int32, y int32) {

	var result *calculatorpb.CalcResult

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("failed to connect to server with err %v", err)
	}
	defer cc.Close()

	cv := calculatorpb.NewVersionServiceClient(cc)

	res, _ := cv.GetVersion(context.Background(), &calculatorpb.VersionRequest{})

	fmt.Println(res.Version)

	c := calculatorpb.NewCalculatorServiceClient(cc)

	req := &calculatorpb.CalcRequest{
		Calc: &calculatorpb.Calc{
			X: x,
			Y: y,
		},
	}

	result, err = c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to connect to server with err %v", err)
	}

	fmt.Printf("Result of %v + %v is: %v", x, y, result.GetValue())

	result, err = c.Multiply(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to connect to server with err %v", err)
	}

	fmt.Printf("Result of %v * %v is: %v", x, y, result.GetValue())
}

func main() {

	gRPCCalculate(10, 3)

}

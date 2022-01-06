package Calculator

import (
	"context"
	"fmt"

	"github.com/apavanello/gRPC-Calculato/pkg/calculatorpb"
)

type Server struct{}

func (s *Server) Sum(ctx context.Context, req *calculatorpb.CalcRequest) (res *calculatorpb.CalcResult, err error) {
	fmt.Println("strat Calculator Sum Service")
	result := req.GetCalc().GetX() + req.GetCalc().GetY()

	res = &calculatorpb.CalcResult{
		Value: result,
	}
	err = nil

	return
}

func (s *Server) Multiply(ctx context.Context, req *calculatorpb.CalcRequest) (res *calculatorpb.CalcResult, err error) {
	fmt.Println("strat Calculator Multiply Service")
	result := req.GetCalc().GetX() * req.GetCalc().GetY()

	res = &calculatorpb.CalcResult{
		Value: result,
	}

	err = nil

	return
}

package calc

import (
	"calculator-grpc-server/api"
	"context"
	"errors"
	"log"
)

type Server struct {
	api.UnimplementedCalculatorServer
}

func (s *Server) DoCalc(ctx context.Context, in *api.CalculatorInputMessage) (*api.CalculatorResultMessage, error) {
	number1 := in.Number1
	number2 := in.Number2
	opr := in.Operator
	log.Printf("Do %s for %d and %d", opr, number1, number2)
	var result int32
	var err error
	switch opr {
	case "+":
		result = number1 + number2
	case "-":
		result = number1 - number2
	case "*":
		result = number1 * number2
	case "/":
		if number2 == 0 {
			err = errors.New("divided by zero")
			return nil, err
		}
		result = number1 / number2
	}

	resultMessage := &api.CalculatorResultMessage{
		ResNumber: result,
	}
	return resultMessage, nil
}

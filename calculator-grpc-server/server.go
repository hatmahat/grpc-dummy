package main

import (
	"calculator-grpc-server/api"
	"calculator-grpc-server/calc"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	//host := os.Getenv("GRPC_HOST")
	//port := os.Getenv("GRPC_PORT")
	host := "localhost"
	port := "8888"
	serverInfo := fmt.Sprintf("%s:%s", host, port)
	listen, err := net.Listen("tcp", serverInfo)
	if err != nil {
		log.Fatalln("failed to listen")
	}

	s := calc.Server{}
	grpcServer := grpc.NewServer()
	api.RegisterCalculatorServer(grpcServer, &s)
	log.Println("Server run on", serverInfo)
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalln("Failed to server...", err)
	}

}

/*
set GRPC_HOST=localhost
set GRPC_PORT=8888
*/

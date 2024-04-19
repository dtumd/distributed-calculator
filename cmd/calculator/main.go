package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	calc "yc/distr-calc/calculator"
	calcp "yc/distr-calc/proto"
)

type Server struct {
	calcp.CalculateServiceServer
}

func NewServer() *Server {
	return &Server{}
}

type CalculateServiceServer interface {
	Calculate(context.Context, *calcp.CalculateRequest) (*calcp.CalculateResponse, error)
	mustEmbedUnimplementedGeometryServiceServer()
}

func (s *Server) Calculate(ctx context.Context, req *calcp.CalculateRequest) (*calcp.CalculateResponse, error) {
	res, err := calc.Calculate(req.GetExpression())
	return &calcp.CalculateResponse{
		Result: res,
	}, err
}

// go run .\cmd\calculator\main.go
func main() {

	fmt.Println("starting")

	host := "localhost"
	port := "8333"

	addr := fmt.Sprintf("%s:%s", host, port)
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Println("error starting tcp listener: ", err)
		os.Exit(1)
	}

	log.Println("tcp listener started at port: ", port)

	grpcServer := grpc.NewServer()

	// new server
	calculateServiceServer := NewServer()

	// register server
	calcp.RegisterCalculateServiceServer(grpcServer, calculateServiceServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Println("error serving grpc: ", err)
		os.Exit(1)
	}

}

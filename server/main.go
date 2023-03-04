package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	// "github.com/go-kit/kit/examples/addsvc/pb"
	pb "github.com/parikalp/grpc4/proto"
	"google.golang.org/grpc"
)

var (
	port = "5050"
	host = "localhost"
)

type Server struct {
	pb.MathServiceServer
}

func main() {
	addr := fmt.Sprintf("%s:%s", host, port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println("tcp listener started at end ")
		os.Exit(1)
	}

	log.Println("tcp listener started at port: ", port)

	grpcServer := grpc.NewServer()
	pb.RegisterMathServiceServer(grpcServer, &Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Println("error serving grpc: ", err)
		os.Exit(1)
	}

}

func (s *Server) Add(ctx context.Context, in *pb.TwoNumRequest) (*pb.TwoNumResponse, error) {
	log.Println("Adding ", in)
	return &pb.TwoNumResponse{
		C: in.A + in.B,
	}, nil
}

func (s *Server) Sub(ctx context.Context, in *pb.TwoNumRequest) (*pb.TwoNumResponse, error) {
	log.Println("Subtract ", in)
	return &pb.TwoNumResponse{
		C: in.A - in.B,
	}, nil
}

func (s *Server) Mul(ctx context.Context, in *pb.TwoNumRequest) (*pb.TwoNumResponse, error) {
	log.Println("Multiply ", in)
	return &pb.TwoNumResponse{
		C: in.A * in.B,
	}, nil
}

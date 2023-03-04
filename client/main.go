package main

import (
	"context"
	"fmt"

	pb "github.com/parikalp/grpc4/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	port = "5050"
	host = "localhost"
)

func main() {
	addr := fmt.Sprintf("%s:%s", host, port)
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("could not connect to grpc server ")
	}

	defer conn.Close()

	grpcClient := pb.NewMathServiceClient(conn)

	add, err := grpcClient.Add(context.TODO(), &pb.TwoNumRequest{
		A: 3.1,
		B: 6.1,
	})
	if err != nil {
		fmt.Println("failed invoking Adding: ", err)
	}

	sub, err := grpcClient.Sub(context.TODO(), &pb.TwoNumRequest{
		A: 5.1,
		B: 2.1,
	})

	if err != nil {
		fmt.Println("failed invoking Subtracting: ", err)
	}

	mul, err := grpcClient.Mul(context.TODO(), &pb.TwoNumRequest{
		A: 5.0,
		B: 2.0,
	})

	if err != nil {
		fmt.Println("failed invoking Multiply: ", err)
	}

	fmt.Println("Adding between two number is ", add.C)
	fmt.Println("Subtracting between two number is ", sub.C)
	fmt.Println("Multiply between two number is ", mul.C)
}

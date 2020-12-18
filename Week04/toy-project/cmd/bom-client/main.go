package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	v1 "toy/api/bom/v1"
)


func main() {
	conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	grpcClient := v1.NewBomClient(conn)
	req := &v1.CreateBomRequest{
		FileName: "bom1.txt",
	}
	grpcClient.CreateBom(context.Background(), req)
}

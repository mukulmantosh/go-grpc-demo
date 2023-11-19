package main

import (
	"context"
	pb "go-grpc-demo/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello"}, nil
}
func main() {

}

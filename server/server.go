package server

import (
	"github.com/grpc/grpc-go"
	"golang.org/x/net/context"
)

type server struct{}

func (s *server) Start(ctx context.Context, in *pb.ProcessStateRequest) (*pb.ProcessStateReply, error) {

}

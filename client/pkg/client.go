package pkg

import (
	"google.golang.org/grpc"

	pb "github.com/cpg1111/spawnd/protobufs"
)

func Connect(addr string) (pb.ProcessClient, error) {
	if addr == "" {
		addr = "/var/run/spawnd.sock"
	}
	conn, connErr := grpc.Dial(addr, grpc.WithInsecure())
	if connErr != nil {
		return nil, connErr
	}
	defer conn.Close()
	cli := pb.NewProcessClient(conn)
	return cli, nil
}

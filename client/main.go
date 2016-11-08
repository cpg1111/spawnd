package client

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"google.golang.org/grpc"

	cli "github.com/cpg1111/spawnd/client/pkg"
	pb "github.com/cpg1111/spawnd/protobufs"
)

func main() {
	payload := &pb.ProcessStateRequest
	strArg := fmt.Sprint(os.Args[2:]...)
	unmarshErr := json.Unmarshal(([]byte)(strArg), payload)
	switch os.Args[1] {
	case "start":
		cli.Start(context.Background(), payload, grpc.WithInsecure())
		break
	case "stop":
		cli.Stop(context.Background(), payload, grpc.WithInsecure())
		break
	case "reload":
		cli.Reload(context.Background(), payload, grpc.WithInsecure())
		break
	case "restart":
		cli.Restart(context.Background(), payload, grpc.WithInsecure())
		break
	}
}

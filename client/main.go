package client

import (
	"encoding/json"
	"log"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	cli "github.com/cpg1111/spawnd/client/pkg"
	pb "github.com/cpg1111/spawnd/protobufs"
)

func concatArgs(args []string) []byte {
	var result []byte
	for i := range args {
		argBytes := ([]byte)(args[i])
		result = append(result, argBytes...)
	}
	return result
}

func main() {
	payload := &pb.ProcessStateRequest{}
	arg := concatArgs(os.Args[2:])
	unmarshErr := json.Unmarshal(arg, payload)
	if unmarshErr != nil {
		log.Fatal(unmarshErr)
	}
	conn, err := cli.Connect(os.Getenv("SPAWND_SERVER_ADDR"))
	if err != nil {
		log.Fatal(err)
	}
	var (
		resp    *pb.ProcessStateReply
		respErr error
	)
	switch os.Args[1] {
	case "start":
		resp, respErr = conn.Start(context.Background(), payload, grpc.FailFast(false))
		break
	case "stop":
		resp, respErr = conn.Stop(context.Background(), payload, grpc.FailFast(false))
		break
	case "reload":
		resp, respErr = conn.Reload(context.Background(), payload, grpc.FailFast(false))
		break
	case "restart":
		resp, respErr = conn.Restart(context.Background(), payload, grpc.FailFast(false))
		break
	}
	if respErr != nil {
		log.Fatal(respErr)
	}
	log.Println(resp)
}

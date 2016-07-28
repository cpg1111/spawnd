package server

import (
	"log"
	"net"

	"github.com/grpc/grpc-go"
	"golang.org/x/net/context"

	"github.com/cpg1111/spawnd/config"
	"github.com/cpg1111/spawnd/daemon"
	pb "github.com/cpg1111/spawnd/protobufs"
)

type Server struct {
	Daemon  *daemon.Daemon
	GRPCSrv *grpc.Server
	TCP     *net.Listener
	pServer *server
}

type server struct {
	daemon *daemon.Daemon
}

func New(conf *config.ConnServer, d *daemon.Daemon) *Server {
	tcp, tcpErr := net.Listen(conf.Type(), conf.Addr())
	if tcpErr != nil {
		log.Fatal(tcpErr)
	}
	grpcSrv := grpc.NewServer()
	srv := &server{
		daemon: d,
	}
	pb.RegisterProcessServer(grpcSrv, srv)
	grpcSrv.Serve(tcp)
	return &Server{
		Daemon:  d,
		GRPCSrv: grpcSrv,
		TCP:     tcp,
		pServer: srv,
	}
}

func getProc(in *pb.ProcessStateRequest) *daemon.Proc {
	if in.PID > 0 {
		proc = s.daemon.GetProc(in.PID)
	} else {
		proc = s.daemon.GetProcByName(in.Name)
	}
}

func (s *server) Start(ctx context.Context, in *pb.ProcessStateRequest) (*pb.ProcessStateReply, error) {
	proc := getProc(in)
	_, err := proc.Start()
	if err != nil {
		return nil, err
	}
	return &pb.ProcessStateReply{
		Name:    in.Name,
		State:   "started",
		Message: "successfully started process",
	}, nil
}

func (s *server) Stop(ctx context.Context, in *pb.ProcessStateRequest) (*pb.ProcessStateReply, error) {
	proc := getProc(in)
	err := proc.Stop()
	if err != nil {
		return nil, err
	}
	return &pb.ProcessStateReply{
		Name:    in.Name,
		State:   "stopped",
		Message: "successfully stopped process",
	}, nil
}

func (s *server) Restart(ctx context.Context, in *pb.ProcessStateRequest) (*pb.ProcessStateReply, error) {
	proc := getProc(in)
	_, err := proc.Restart()
	if err != nil {
		return nil, err
	}
	return &pb.ProcessStateReply{
		Name:    in.Name,
		State:   "started",
		Message: "successfully restarted process",
	}, nil
}

func (s *server) Reload(ctx context.Context, in *pb.ProcessStateRequest) (*pb.ProcessStateReply, error) {
	proc := getProc(in)
	err := proc.Reload()
	if err != nil {
		return nil, err
	}
	return &pb.ProcessStateReply{
		Name:    in.Name,
		State:   "started",
		Message: "successfuly reloaded process",
	}, nil
}

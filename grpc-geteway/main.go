package main

import (
	"context"
	"flag"
	"fmt"
	v1 "gateway/service/v1"
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
)

type server struct {
	v1.UnimplementedYourServiceServer
}

func (s *server) Echo(ctx context.Context, msg *v1.StringMessage) (*v1.StringMessage, error) {
	msg.Value = msg.Value + "ffsdffds"
	fmt.Println(123123123)
	return msg, nil
}

func run() error {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	v1.RegisterYourServiceServer(s, &server{})
	// Serve gRPC Server
	log.Println("Serving gRPC on 0.0.0.0:8088")
	log.Fatal(s.Serve(lis))
	return err
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}

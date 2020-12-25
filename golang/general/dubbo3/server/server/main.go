package main

import (
	"fmt"
	pb "github.com/apache/dubbo-samples/golang/general/dubbo3/protobuf/grpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	port = ":20000"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) Dubbo3SayHello2(ctx context.Context, in *pb.Dubbo3HelloRequest) (*pb.Dubbo3HelloReply, error) {
	fmt.Println("######### get server request name :" + in.Myname)
	return &pb.Dubbo3HelloReply{Msg: "Hello " + in.Myname}, nil

}

func (s *server) Dubbo3SayHello(svr pb.Dubbo3Greeter_Dubbo3SayHelloServer) error {
	c, err := svr.Recv()
	if err != nil {
		return err
	}
	fmt.Println("server server recv 1 = ", c)
	c2, err := svr.Recv()
	if err != nil {
		return err
	}
	fmt.Println("server server recv 2 = ", c2)
	c3, err := svr.Recv()
	fmt.Println("server server recv 3 = ", c3)

	svr.Send(&pb.Dubbo3HelloReply{
		Msg: c.Myname + c2.Myname,
	})
	fmt.Println("server server send 1 = ", c.Myname+c2.Myname)
	svr.Send(&pb.Dubbo3HelloReply{
		Msg: c3.Myname,
	})
	fmt.Println("server server send 2 = ", c3.Myname)
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDubbo3GreeterServer(s, &server{})
	// Register reflection service on gRPC grpc-grpc-client.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

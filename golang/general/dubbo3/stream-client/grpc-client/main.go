package main

import (
	"fmt"
	pb "github.com/apache/dubbo-samples/golang/general/dubbo3/protobuf/grpc"
	"log"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:20000"
	//address     = "localhost:50051"
)

func main() {
	// Set up a connection to the grpc-client.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDubbo3GreeterClient(conn)

	// Contact the grpc-client and print out its response.
	name := "jifeng"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	for i := 0; i < 100; i++ {
		r, err := c.Dubbo3SayHello(context.Background())
		if err != nil {
			fmt.Println("say hello err:", err)
		}

		if err := r.Send(&pb.Dubbo3HelloRequest{Myname: name}); err != nil {
			fmt.Println("say hello err:", err)
		}
		if err := r.Send(&pb.Dubbo3HelloRequest{Myname: name}); err != nil {
			fmt.Println("say hello err:", err)
		}
		//time.Sleep(time.Second * 10)
		if err := r.Send(&pb.Dubbo3HelloRequest{Myname: name}); err != nil {
			fmt.Println("say hello err:", err)
		}
		rsp := &pb.Dubbo3HelloReply{}
		if err := r.RecvMsg(rsp); err != nil {
			fmt.Println("err = ", err)
		}
		fmt.Printf("firstSend Got rsp = %+v\n", rsp)
		rsp = &pb.Dubbo3HelloReply{}
		if err := r.RecvMsg(rsp); err != nil {
			fmt.Println("err = ", err)
		}
		fmt.Printf("firstSend Got rsp = %+v\n", rsp)

	}

}

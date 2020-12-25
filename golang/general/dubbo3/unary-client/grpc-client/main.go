package main

import (
	pb "github.com/apache/dubbo-samples/golang/general/dubbo3/protobuf/grpc"
	"log"
	"os"
	"sync"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:20000"
	defaultName = "jifeng"
)

func main() {
	// Set up a connection to the grpc-grpc-client.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDubbo3GreeterClient(conn)

	// Contact the grpc-grpc-client and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			r, err := c.Dubbo3SayHello2(context.Background(), &pb.Dubbo3HelloRequest{Myname: name})
			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}
			log.Printf("####### get grpc-grpc-client %+v", r)
			wg.Done()
		}()
	}
	wg.Wait()
}

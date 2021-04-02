package main

import (
	"log"
	"net"

	"github.com/zapkub/bootstore-angular-go-grpc/internal"
	"github.com/zapkub/bootstore-angular-go-grpc/pkg/apis"
	"google.golang.org/grpc"
)

func main() {

	var err error
	var grpcserv = grpc.NewServer()
	apis.RegisterBookstoreServiceServer(grpcserv, internal.BookStoreEndpointHandler{})

	var lis net.Listener
	lis, err = net.Listen("tcp", "127.0.0.1:3000")
	if err != nil {
		panic(err)
	}

	log.Println("serve grpc server at 127.0.0.1:3000")
	err = grpcserv.Serve(lis)
	if err != nil {
		panic(err)
	}

}

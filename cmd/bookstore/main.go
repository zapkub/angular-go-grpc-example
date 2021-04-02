package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"github.com/zapkub/bootstore-angular-go-grpc/internal"
	"github.com/zapkub/bootstore-angular-go-grpc/pkg/apis"
	"google.golang.org/grpc"
)

func SinglePort() {
	var mux = runtime.NewServeMux()

	var ctx = context.Background()
	apis.RegisterBookstoreHandlerServer(ctx, mux, &internal.BookStoreEndpointHandler{})

	var err = http.ListenAndServe("127.0.0.1:8080", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		cors.AllowAll().Handler(mux).ServeHTTP(rw, r)
	}))
	if err != nil {
		panic(err)
	}
}

func MultiplePort() {

	var (
		err             error
		grpcserver      *grpc.Server
		lis             net.Listener
		grpcEndpointURL = "127.0.0.1:3000"
	)

	go func() {
		grpcserver = grpc.NewServer()
		lis, err = net.Listen("tcp", grpcEndpointURL)
		if err != nil {
			panic(err)
		}
		apis.RegisterBookstoreServer(grpcserver, &internal.BookStoreEndpointHandler{})
		log.Println("grpc server listen at :3000")
		err = grpcserver.Serve(lis) // blocking
		if err != nil {
			panic(err)
		}
	}()

	// Gateway
	var mux = runtime.NewServeMux()
	var ctx = context.Background()
	err = apis.RegisterBookstoreHandlerFromEndpoint(
		ctx, mux,
		grpcEndpointURL,
		[]grpc.DialOption{
			grpc.WithInsecure(),
		},
	)
	if err != nil {
		panic(err)
	}
	log.Println("grpc server listen at :8080")
	err = http.ListenAndServe("127.0.0.1:8080", mux)
	if err != nil {
		panic(err)
	}

}

func main() {
	// go MultiplePort()
	go SinglePort()
	var sigch = make(chan os.Signal, 1)
	signal.Notify(sigch, syscall.SIGINT, syscall.SIGTERM)
	<-sigch
	log.Println("close server")

}

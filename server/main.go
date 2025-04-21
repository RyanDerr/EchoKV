package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	service "github.com/RyanDerr/EchoKV/pkg/service/api"
	pb "github.com/RyanDerr/EchoKV/proto-public/go"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	grpcAddr = "localhost:50051"
	httpAddr = "localhost:8080"
)

func main() {
	// Start gRPC server in a separate goroutine
	go func() {
		lis, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			log.Fatalf("Failed to listen: %v\n", err)
		}
		log.Printf("gRPC server listening at %v\n", lis.Addr())

		s := grpc.NewServer()
		svc := service.NewService()
		pb.RegisterEchoKVServer(s, svc)

		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v\n", err)
		}
	}()

	// Create a new gRPC Gateway mux for the HTTP server on main goroutine
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterEchoKVHandlerFromEndpoint(context.Background(), mux, grpcAddr, opts)
	if err != nil {
		log.Fatalf("Failed to start HTTP gateway: %v", err)
	}

	// Start the HTTP server
	log.Printf("Starting HTTP server on %s\n", httpAddr)

	server := configHttpServer(httpAddr, mux)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to serve HTTP: %v", err)
	}
}

func configHttpServer(addr string, mux *runtime.ServeMux) *http.Server {
	return &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

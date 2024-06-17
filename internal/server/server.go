package server

import (
	"log"
	"net"

	"github.com/grantchen2003/insight/summarizer/internal/handlers"
	"github.com/grantchen2003/insight/summarizer/internal/protobufs"
	"google.golang.org/grpc"
)

type Server struct {
	grpcServer *grpc.Server
}

func NewServer() *Server {
	grpcServer := grpc.NewServer()

	protobufs.RegisterSummarizerServiceServer(
		grpcServer, &handlers.SummarizerServiceHandler{},
	)

	return &Server{grpcServer: grpcServer}
}

func (server Server) Start(address string) error {
	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}

	log.Printf("server listening on %s", address)

	if err := server.grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return err
	}

	return nil
}

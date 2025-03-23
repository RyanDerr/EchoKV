package service

import (
	pb "github.com/RyanDerr/EchoKV/proto-public/go"
)

// Service implements the gRPC service "EchoKV".
type Service struct {
	pb.UnimplementedEchoKVServer
	cache map[string]string
}

// NewService creates a new Service.
func NewService() *Service {
	return &Service{
		cache: make(map[string]string),
	}
}

// Ensure that Service implements the EchoKV interface.
var (
	_ pb.EchoKVServer = (*Service)(nil)
)

package service

import (
	pb "github.com/RyanDerr/GoKeyValueStore/proto-public/go"
)

// Service implements the gRPC service "CacheService".
type Service struct {
	pb.UnimplementedKeyValueServiceServer
	cache map[string]string
}

// NewService creates a new Service.
func NewService() *Service {
	return &Service{
		cache: make(map[string]string),
	}
}

// Ensure that Service implements the CacheServiceServer interface.
var (
	_ pb.KeyValueServiceServer = (*Service)(nil)
)

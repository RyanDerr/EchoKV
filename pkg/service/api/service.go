package service

import (
	"log/slog"
	"os"
	"sync"

	pb "github.com/RyanDerr/EchoKV/proto-public/go"
	"github.com/hashicorp/raft"
)

// Service implements the gRPC service "EchoKV".
type Service struct {
	pb.UnimplementedEchoKVServer
	store  map[string]string
	raft   *raft.Raft
	mu     sync.Mutex
	logger *slog.Logger
}

// NewService creates a new Service.
func NewService() *Service {
	return &Service{
		store:  make(map[string]string),
		logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}
}

// Ensure that Service implements the EchoKV interface.
var (
	_ pb.EchoKVServer = (*Service)(nil)
)

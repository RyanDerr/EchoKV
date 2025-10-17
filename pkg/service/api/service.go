package service

import (
	"io"
	"log/slog"
	"os"

	pb "github.com/RyanDerr/EchoKV/proto-public/go"
	"github.com/hashicorp/raft"
)

// Service implements the gRPC service "EchoKV".
type Service struct {
	pb.UnimplementedKeyValueServer
	store map[string]string
	// raft   *raft.Raft
	// mu     sync.Mutex
	logger *slog.Logger
}

// NewService creates a new Service.
func NewService() *Service {
	return &Service{
		store:  make(map[string]string),
		logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}
}

func (s *Service) Open(dataDir, bindAddr, singleNode string) error {
	// Initialize the Raft configuration and start the Raft node.
	// This is a placeholder for the actual Raft initialization code.
	return nil
}

func (s *Service) Join(addr string) error {
	// Join a Raft cluster.
	// This is a placeholder for the actual Raft join code.
	return nil
}

func (s *Service) Apply(log *raft.Log) any {
	// Apply a log entry to the state machine.
	// This is a placeholder for the actual Raft apply code.
	return nil
}

func (s *Service) Snapshot() (raft.FSMSnapshot, error) {
	// Create a snapshot of the state machine.
	// This is a placeholder for the actual Raft snapshot code.
	return nil, nil
}

func (s *Service) Restore(snap io.ReadCloser) error {
	// Restore the state machine from a snapshot.
	// This is a placeholder for the actual Raft restore code.
	return nil
}

func (s *Service) Persist(snap raft.SnapshotSink) error {
	// Persist the snapshot to disk.
	// This is a placeholder for the actual Raft persist code.
	return nil
}

func (s *Service) Release() {
	// Release any resources held by the service.
	// This is a placeholder for the actual resource release code.
}

// Ensure that Service implements the EchoKV interface.
var (
	_ pb.KeyValueServer = (*Service)(nil)
)

package controlplane

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/google/uuid"
)

var (
	logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
)

// Cluster is a collection of nodes
type Cluster struct {
	Name string
	ID   uuid.UUID
	Pool map[uuid.UUID]*Node
}

type ClusterOptions struct {
	Name string
	Size int
}

// DefaultClusterSize is 3, which allows for a quorum of 2 writes
const DefaultClusterSize = 3

// NewCluster creates a new cluster
func NewCluster(options ...ClusterOptions) (*Cluster, error) {

	id := uuid.New()

	logger.Info("Creating a new cluster", "clusterID", id)
	c := &Cluster{
		Pool: make(map[uuid.UUID]*Node),
		ID:   id,
	}

	var opts ClusterOptions
	if len(options) > 0 {
		opts = options[0]
	}

	clusterSize := opts.Size
	if clusterSize < DefaultClusterSize {
		logger.Warn("Invalid cluster size, using default", "clusterSize", DefaultClusterSize)
		clusterSize = DefaultClusterSize
	}

	name := opts.Name
	if name == "" {
		logger.Debug("No cluster name provided, generating a new one")
		name = generateRandomClusterName(id)
	}

	c.Name = name

	for range clusterSize {
		//Create node and add node
		if err := c.AddNode(); err != nil {
			logger.Error("Failed to add node", "error", err)
			return nil, err
		}
	}

	return c, nil
}

// AddNode adds a new node to the cluster
func (c *Cluster) AddNode() error {
	node := NewNode(c.ID)

	// This should never happen in practice, but it's good to check
	if _, exists := c.Pool[node.ID]; exists {
		logger.Warn("Node already exists", "name", node.Name)
		return fmt.Errorf("node with name %s already exists", node.Name)
	}

	c.Pool[node.ID] = node
	logger.Info("Node added successfully", "name", node.Name)
	return nil
}

func generateRandomClusterName(ID uuid.UUID) string {
	return fmt.Sprintf("cluster-%s", ID.String())
}

// GetClusterSize returns the size of the cluster
func (c *Cluster) GetClusterSize() int {
	return len(c.Pool)
}

// GetClusterID returns the ID of the cluster
func (c *Cluster) GetClusterID() uuid.UUID {
	return c.ID
}

// GetClusterName returns the name of the cluster
func (c *Cluster) GetClusterName() string {
	return c.Name
}

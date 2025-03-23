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
	ID   uuid.UUID
	Pool map[string]*Node
}

type ClusterOptions struct {
	Size int
}

// DefaultClusterSize is 3, which allows for a quorum of 2 writes
const DefaultClusterSize = 3

// NewCluster creates a new cluster
func NewCluster(options ...ClusterOptions) (*Cluster, error) {

	clusterID := uuid.New()

	logger.Info("Creating a new cluster", "clusterID", clusterID)
	cluster := &Cluster{
		Pool: make(map[string]*Node),
		ID:   clusterID,
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

	for i := 0; i < clusterSize; i++ {
		//Create node and add node
		if err := cluster.AddNode(); err != nil {
			logger.Error("Failed to add node", "error", err)
			return nil, err
		}

		logger.Debug("Successfully added node to cluster", "nodeCount", i+1)
	}

	return cluster, nil
}

// GetClusterSize returns the size of the cluster
func (c *Cluster) GetClusterSize() int {
	return len(c.Pool)
}

func (c *Cluster) AddNode() error {
	node := NewNode(c.ID)

	//This should never happen in practice, but it's good to check
	if node == nil {
		logger.Error("Failed to create node")
		return fmt.Errorf("failed to create node")
	}

	// This should never happen in practice, but it's good to check
	if _, exists := c.Pool[node.Name]; exists {
		logger.Warn("Node already exists", "name", node.Name)
		return fmt.Errorf("node with name %s already exists", node.Name)
	}

	c.addNode(node)
	logger.Info("Node added successfully", "name", node.Name)
	return nil
}

// addNode adds a node to the internal pool by name (internal use only).
func (c *Cluster) addNode(node *Node) {
	c.Pool[node.Name] = node
}

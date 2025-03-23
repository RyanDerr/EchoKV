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
	Pool map[string]*Node
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
	cluster := &Cluster{
		Pool: make(map[string]*Node),
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
		logger.Info("No cluster name provided, generating a new one")
		name = generateRandomClusterName(id)
	}

	cluster.Name = name

	for range clusterSize {
		//Create node and add node
		if err := cluster.AddNode(); err != nil {
			logger.Error("Failed to add node", "error", err)
			return nil, err
		}
	}

	return cluster, nil
}

// GetClusterSize returns the size of the cluster
func (c *Cluster) GetClusterSize() int {
	return len(c.Pool)
}

func (c *Cluster) AddNode() error {
	node := NewNode(c.ID)

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

func generateRandomClusterName(uniqueID uuid.UUID) string {
	return fmt.Sprintf("cluster-%s", uniqueID.String())
}

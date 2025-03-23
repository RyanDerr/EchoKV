package controlplane

import (
	"fmt"

	"github.com/google/uuid"
)

type Node struct {
	Name      string
	UniqueID  uuid.UUID
	ClusterID uuid.UUID
}

// NodeOptions is a set of optional parameters for creating a new node
type NodeOptions struct {
	Name string
}

func NewNode(clusterID uuid.UUID, options ...NodeOptions) *Node {
	id := uuid.New()
	logger.Info("Creating a new node", "clusterID", clusterID)
	logger.Debug("Generated Node ID", "id", id)

	var opts NodeOptionsß
	if len(options) > 0 {
		opts = options[0]
	}

	name := opts.Name
	if name == "" {
		logger.Info("No name provided, generating a new one")
		name = generateName(id)
		logger.Info("Generated name", "name", name)
	}

	return &Node{
		Name:      name,
		UniqueID:  id,
		ClusterID: clusterID,
	}
}

func generateName(uniqueID uuid.UUID) string {
	return fmt.Sprintf("node-%s", uniqueID.String())
}

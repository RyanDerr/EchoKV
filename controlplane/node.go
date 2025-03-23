package controlplane

import (
	"fmt"

	"github.com/google/uuid"
)

type Node struct {
	Name      string
	ID        uuid.UUID
	ClusterID uuid.UUID
}

func NewNode(clusterID uuid.UUID) *Node {
	id := uuid.New()
	logger.Info("Creating a new node", "clusterID", clusterID)
	logger.Debug("Generated Node ID", "id", id)

	name := generateRandomNodeName(id)
	logger.Debug("Generated name", "name", name)

	return &Node{
		Name:      name,
		ID:        id,
		ClusterID: clusterID,
	}
}

func generateRandomNodeName(ID uuid.UUID) string {
	return fmt.Sprintf("node-%s", ID.String())
}

package controlplane

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewNode_GeneratesValidNode(t *testing.T) {
	clusterID := uuid.New()
	n := NewNode(clusterID)

	assert.NotNil(t, n, "node should not be nil")
	assert.Equal(t, clusterID, n.ClusterID, "cluster ID should be assigned")

	assert.NotEmpty(t, n.ID, "node ID should be generated")
	assert.NotEmpty(t, n.Name, "node name should be generated")

	assert.True(t, strings.HasPrefix(n.Name, "node-"), "node name should start with 'node-'")
}

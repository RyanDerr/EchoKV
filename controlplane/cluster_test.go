package controlplane

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCluster_DefaultsApplied(t *testing.T) {

	c, err := NewCluster()

	assert.NoError(t, err)
	assert.Equal(t, DefaultClusterSize, c.GetClusterSize(), fmt.Sprintf("Expected default cluster pool size to be %d", DefaultClusterSize))
	assert.True(t, strings.HasPrefix(c.Name, "cluster-"), "Expected default cluster name to be prefixed with  'cluster-'")

}

func TestNewCluster_WithCustomOptions(t *testing.T) {
	customSize := 5
	customName := "test-cluster"

	c, err := NewCluster(ClusterOptions{
		Name: customName,
		Size: customSize,
	})

	assert.NoError(t, err)
	assert.Equal(t, customSize, c.GetClusterSize(), "Expected custom cluster size to be applied")
	assert.Equal(t, customName, c.Name, "Expected custom cluster name to be applied")
}

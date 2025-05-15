package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeReplication(t *testing.T) {
	node := NewBaseNode("TestNodeType", "testId")
	replication := node.GetReplication()

	node.Id = "testId2"

	ds := BuildChangeSet(replication)
	assert.Equal(t, 1, len(ds))
	dsNode := ds["node"].(map[string]any)
	assert.Equal(t, 1, len(dsNode))
	assert.Equal(t, "testId2", dsNode["id"])

	node.nodeType = "testNodeType2"
	ds = BuildChangeSet(replication)
	assert.Equal(t, 1, len(ds))
	dsNode = ds["node"].(map[string]any)
	assert.Equal(t, 2, len(dsNode))
	assert.Equal(t, "testId2", dsNode["id"])
	assert.Equal(t, "testNodeType2", dsNode["type"])

	node.Transform.Position.X = 69
	ds = BuildChangeSet(replication)
	assert.Equal(t, 1, len(ds))
	dsNode = ds["node"].(map[string]any)
	assert.Equal(t, 3, len(dsNode))
	assert.Equal(t, "testId2", dsNode["id"])
	assert.Equal(t, "testNodeType2", dsNode["type"])
	dsTransform := dsNode["transform"].(map[string]any)
	assert.Equal(t, float32(69), dsTransform["position.x"])
}

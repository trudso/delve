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
	dsNode := ds["testId"].(map[string]any)
	assert.Equal(t, 1, len(dsNode))
	assert.Equal(t, "testId2", dsNode["id"])

	node.nodeType = "TestNodeType2"
	ds = BuildChangeSet(replication)
	assert.Equal(t, 1, len(ds))
	dsNode = ds["testId"].(map[string]any)
	assert.Equal(t, 2, len(dsNode))
	assert.Equal(t, "testId2", dsNode["id"])
	assert.Equal(t, "TestNodeType2", dsNode["type"])

	node.Transform.Position.X = 69
	ds = BuildChangeSet(replication)
	assert.Equal(t, 1, len(ds))
	dsNode = ds["testId"].(map[string]any)
	assert.Equal(t, 3, len(dsNode))
	assert.Equal(t, "testId2", dsNode["id"])
	assert.Equal(t, "TestNodeType2", dsNode["type"])
	dsTransform := dsNode["transform"].(map[string]any)
	assert.Equal(t, float32(69), dsTransform["position.x"])

	replication.ResetToChanged()
	ds = BuildChangeSet(replication)
	assert.Empty(t, ds)
} 

func TestNodeWithChildrenApplyDataSet(t *testing.T) {
	node := NewBaseNode("TestNodeType", "testId")

	//node.Id = "testId2"
	node.Transform.Position.X = 0
	node.Transform.Position.Y = 1
	node.Transform.Scale.X = 2
	node.Transform.Scale.Y = 3
	node.Transform.Rotation.X = 4
	node.Transform.Rotation.Y = 5

	c1 := NewBaseNode("ChildNode", "child1")
	c1.Transform.Position.X = 1
	node.AddChild(&c1) 

	c2 := NewBaseNode("ChildNode", "child2")
	c2.Transform.Position.X = 2
	node.AddChild(&c2) 
	
	// since it's not a changeset we can just build the replication here
	rep := node.GetReplication()
	ds := BuildSnapshot(rep)

	newNode := NewBaseNode("_", "testId" )
	newReplication := newNode.GetReplication()

	ApplyDataSet(newReplication, ds)
	assert.Equal(t, node, newNode)
}

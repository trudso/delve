package engine

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func baseNodeFactory(id string, _ map[string]any) Node {
	node := NewBaseNode(id, reflect.TypeOf(BaseNode{}), baseNodeFactory)
	return &node
}

func TestNodeReplication(t *testing.T) {
	node := NewBaseNode("testId", reflect.TypeOf(BaseNode{}), nil)
	replication := node.GetReplication()
	node.Transform.Position.X = 69

	ds := BuildChangeSet(replication)
	assert.Equal(t, 1, len(ds))
	dsNode := ds["testId"].(map[string]any)
	assert.Equal(t, 3, len(dsNode))
	assert.Equal(t, "testId", dsNode["id"])
	assert.Equal(t, "engine.BaseNode", dsNode["type"])
	dsTransform := dsNode["transform"].(map[string]any)
	assert.Equal(t, 1, len(dsTransform))
	assert.Equal(t, float32(69), dsTransform["position.x"])

	node.Transform.Position.Y = 420
	ds = BuildChangeSet(replication)
	assert.Equal(t, 1, len(ds))
	dsNode = ds["testId"].(map[string]any)
	assert.Equal(t, 3, len(dsNode))
	assert.Equal(t, "testId", dsNode["id"])
	assert.Equal(t, "engine.BaseNode", dsNode["type"])
	dsTransform = dsNode["transform"].(map[string]any)
	assert.Equal(t, 2, len(dsTransform))
	assert.Equal(t, float32(69), dsTransform["position.x"])
	assert.Equal(t, float32(420), dsTransform["position.y"])

	replication.ResetToChanged()
	ds = BuildChangeSet(replication)
	assert.Empty(t, ds)
}

func TestNodeWithChildrenApplySnapshot(t *testing.T) {
	baseNodeType := reflect.TypeOf(BaseNode{})
	node := NewBaseNode("testId", baseNodeType, baseNodeFactory)

	node.Transform.Position.X = 0
	node.Transform.Position.Y = 1
	node.Transform.Scale.X = 2
	node.Transform.Scale.Y = 3
	node.Transform.Rotation.X = 4
	node.Transform.Rotation.Y = 5

	c1 := NewBaseNode("child1", baseNodeType, baseNodeFactory)
	c1.Transform.Position.X = 1
	node.AddChild(&c1)

	c2 := NewBaseNode("child2", baseNodeType, baseNodeFactory)
	c2.Transform.Position.X = 2
	node.AddChild(&c2)

	// since it's not a changeset we can just build the replication here
	rep := node.GetReplication()
	ds := BuildSnapshot(rep)

	//SaveJson("/home/trudso/Dev/Delve.go/src", "test.json", ds)

	newNode := NewBaseNode("testId", baseNodeType, baseNodeFactory)
	newReplication := newNode.GetReplication()

	ApplyDataSet(newReplication, ds)
	//assert.Equal(t, node, newNode)

	assert.Equal(t, 2, len(newNode.GetChildren()))

	nc1 := newNode.GetChild("child1")
	nc2 := newNode.GetChild("child2")
	assert.Equal(t, float32(1), nc1.GetTransform().Position.X)
	assert.Equal(t, float32(2), nc2.GetTransform().Position.X)
}

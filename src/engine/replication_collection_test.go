package engine

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransformCollectionReplication(t *testing.T) {
	transform := NewTransform()

	serverAuthority := func() bool { return true }
	replication := NewReplicationCollection("transform", []Replicatable{
		NewReplicationPrimitive("position.x", &transform.Position.X, true, serverAuthority),
		NewReplicationPrimitive("position.y", &transform.Position.Y, true, serverAuthority),
	}, nil, nil)

	// add changes
	transform.Position.X = 42
	ds := BuildChangeSet(replication)

	assert.Equal(t, 1, len(ds))
	posData := ds["transform"].(map[string]any)
	assert.Equal(t, 1, len(posData))
	assert.Equal(t, float32(42), posData["position.x"])

	// change again
	transform.Position.X = 69
	transform.Position.Y = 420

	// build change set
	ds = BuildChangeSet(replication)
	assert.Equal(t, 1, len(ds))
	posData = ds["transform"].(map[string]any)
	assert.Equal(t, 2, len(posData))
	assert.Equal(t, float32(69), posData["position.x"])
	assert.Equal(t, float32(420), posData["position.y"])

	// reset changed flag
	ResetToChanged(replication)
	ds = BuildChangeSet(replication)
	assert.Equal(t, 0, len(ds))

	// fetch snapshot
	ds = BuildSnapshot(replication)
	assert.Equal(t, 1, len(ds))
	posData = ds["transform"].(map[string]any)
	assert.Equal(t, 2, len(posData))
	assert.Equal(t, float32(69), posData["position.x"])
	assert.Equal(t, float32(420), posData["position.y"])
}

func TestChildReplication(t *testing.T) {
	baseNodeType := reflect.TypeOf(BaseNode{})
	node := NewBaseNode("testId", baseNodeType, baseNodeFactory)

	node.Transform.Position.X = 0
	node.Transform.Position.Y = 1
	node.Transform.Scale.X = 2
	node.Transform.Scale.Y = 3
	node.Transform.Rotation.X = 4
	node.Transform.Rotation.Y = 5

	c1 := NewBaseNode("child1", baseNodeType, baseNodeFactory)
	node.AddChild(&c1)

	// since it's not a changeset we can just build the replication here
	rep := node.GetReplication()
	c1.Transform.Position.Y = 1
	ds := BuildChangeSet(rep)

	//SaveJson("/home/trudso/Dev/Delve.go/src", "test.json", ds)
	newNode := NewBaseNode("testId", baseNodeType, baseNodeFactory)
	newReplication := newNode.GetReplication()

	ApplyDataSet(&newReplication, ds)
	//assert.Equal(t, node, newNode)

	assert.Equal(t, 1, len(newNode.GetChildren()))

	nc1 := newNode.GetChild("child1")
	assert.Equal(t, float32(1), nc1.GetTransform().Position.Y)
}

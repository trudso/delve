package engine

import (
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

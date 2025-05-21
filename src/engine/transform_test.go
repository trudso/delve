package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranformReplication(t *testing.T) {
	transform := NewTransform()
	transform.Position.X = 69
	transform.Position.Y = 420
	transform.Scale.X = 1
	transform.Scale.Y = 2
	transform.Rotation.X = 3
	transform.Rotation.Y = 4

	replication := transform.GetReplication()

	ds := BuildSnapshot(replication)
	assert.Equal(t, 1, len(ds))
	dsT := ds["transform"].(map[string]any)
	assert.Equal(t, 6, len(dsT))
	assert.Equal(t, float32(69), dsT["position.x"])
	assert.Equal(t, float32(420), dsT["position.y"])
	assert.Equal(t, float32(1), dsT["scale.x"])
	assert.Equal(t, float32(2), dsT["scale.y"])
	assert.Equal(t, float32(3), dsT["rotation.x"])
	assert.Equal(t, float32(4), dsT["rotation.y"])

	// test changesets
	ResetToChanged(replication)
	ds = BuildChangeSet(replication)
	assert.Empty(t, ds) 

	// test changing
	transform.Position.X = 2 
	transform.Position.Y = 29 

	ds = BuildChangeSet(replication)
	assert.Equal(t, 1, len(ds))
	dsT = ds["transform"].(map[string]any)
	assert.Equal(t, 2, len(dsT))
	assert.Equal(t, float32(2), dsT["position.x"])
	assert.Equal(t, float32(29), dsT["position.y"])
}

func TestTransformApplyDataSet(t *testing.T) {
	transform := NewTransform()
	transform.Position.X = 1
	transform.Position.Y = 2
	transform.Scale.X = 3 
	transform.Scale.Y = 4
	transform.Rotation.X = 5
	transform.Rotation.Y = 6

	replication := transform.GetReplication()
	ds := BuildSnapshot(replication)

	newTransform := NewTransform()
	newReplication := newTransform.GetReplication()

	ApplyDataSet(newReplication, ds)
	assert.Equal(t, transform, newTransform)
}

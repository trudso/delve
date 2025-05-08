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

	ds := BuildSnapshot(transform.GetReplication())
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
	ResetToChanged(transform.GetReplication())
	ds = BuildChangeSet(transform.GetReplication())
	assert.Equal(t, 0, len(ds))

	// test changing
	transform.Position.X = 2 
	transform.Position.Y = 29 

	ds = BuildChangeSet(transform.GetReplication())
	assert.Equal(t, 1, len(ds))
	dsT = ds["transform"].(map[string]any)
	assert.Equal(t, 2, len(dsT))
	assert.Equal(t, float32(2), dsT["position.x"])
	assert.Equal(t, float32(29), dsT["position.y"])
}

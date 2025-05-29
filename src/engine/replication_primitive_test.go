package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransformPrimitiveReplication(t *testing.T) {
	transform := NewTransform()

	serverAuthority := func() bool { return true }
	replication := NewReplicationPrimitive("position.x", &transform.Position.X, true, serverAuthority)

	// add changes
	transform.Position.X = 42
	transform.Position.Y = 69
	ds := BuildChangeSet(replication)

	assert.Equal(t, 1, len(ds))
	assert.Equal(t, float32(42), ds["position.x"])

	replication.ResetToChanged()
	ds = BuildChangeSet(replication)
	assert.Empty(t, ds)

	// build change set
	transform.Position.X = 420
	ds = BuildChangeSet(replication)
	assert.Equal(t, 1, len(ds))
	assert.Equal(t, float32(420), ds["position.x"])
}

package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransformDataSets(t *testing.T) {
	t1 := NewTransform()

	t1.Position.X = 1.
	t1.Position.Y = 2.
	t1.Rotation.X = 3.
	t1.Rotation.Y = 4.
	t1.Scale.X = 5.
	t1.Scale.Y = 6.

	ds := t1.GetDataSet(false)
	assert.Equal(t, float32(1), ds["position.x"])
	assert.Equal(t, float32(2), ds["position.y"])
	assert.Equal(t, float32(3), ds["rotation.x"])
	assert.Equal(t, float32(4), ds["rotation.y"])
	assert.Equal(t, float32(5), ds["scale.x"])
	assert.Equal(t, float32(6), ds["scale.y"])

	t2 := NewTransform().ApplyDataSet( ds )
	assert.Equal(t, t1.Position.X, t2.Position.X)
	assert.Equal(t, t1.Position.Y, t2.Position.Y)
	assert.Equal(t, t1.Rotation.X, t2.Rotation.X)
	assert.Equal(t, t1.Rotation.Y, t2.Rotation.Y)
	assert.Equal(t, t1.Scale.X, t2.Scale.X)
	assert.Equal(t, t1.Scale.Y, t2.Scale.Y)
}

func TestTransformReplicationToClient(t *testing.T) {
	server := NewReplicatable( NewTransform(), true, func() bool {
		return true
	})

	client := NewReplicatable( NewTransform(), true, func() bool {
		return false
	})

	assert.False( t, server.IsChanged())
	assert.False( t, client.IsChanged())

	// change server transform
	server.Get().Position.X = 69
	server.Get().Position.Y = 420
	assert.True( t, server.IsChanged())
	
	// get data set
	ds := server.Get().GetDataSet( true )

	// apply to client
	clientTransform := client.Get().ApplyDataSet( ds )
	client.SetFromAuthority( clientTransform )

	assert.False( t, client.IsChanged())
	assert.Equal( t, float32(69), client.Get().Position.X )
	assert.Equal( t, float32(420), client.Get().Position.Y )
}

func TestTransformAttemptReplicationToAuthority(t *testing.T) {
	server := NewReplicatable( NewTransform(), true, func() bool {
		return true
	})

	server.Get().Position.X = 1

	ds := server.Get().GetDataSet(true)
	assert.Equal(t, 6, len(ds))
	assert.Equal(t, float32(1), ds["position.x"])

	// hack the data set to attempt to alter server values
	ds["position.x"] = 42.

	changedTransform := server.Get().ApplyDataSet(ds)
	assert.Equal( t, float32(42), changedTransform.Position.X )

	// assert that set from authority on authority does nothing
	server.SetFromAuthority(changedTransform)
	assert.Equal( t, float32(1), server.Get().Position.X )
}

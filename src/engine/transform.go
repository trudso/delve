package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Transform struct {
	Position    rl.Vector2
	Scale       rl.Vector2
	Rotation    rl.Vector2
}

func NewTransform() Transform {
	return Transform{
		Position: rl.Vector2{},
		Scale:    rl.Vector2{X: 1, Y: 1},
		Rotation: rl.Vector2{},
	}
}

func (t *Transform) GetReplication() *ReplicationCollection {
	replication := NewReplicationCollection("transform", []Replicatable{
		NewReplicationPrimitive("position.x", &t.Position.X, true, nil),
		NewReplicationPrimitive("position.y", &t.Position.Y, true, nil),
		NewReplicationPrimitive("scale.x", &t.Scale.X, true, nil),
		NewReplicationPrimitive("scale.y", &t.Scale.Y, true, nil),
		NewReplicationPrimitive("rotation.x", &t.Rotation.X, true, nil),
		NewReplicationPrimitive("rotation.y", &t.Rotation.Y, true, nil),
	}, nil, nil)

	return &replication
}

func getFloat32(number any) float32 {
	n32, ok := number.(float32)
	if ok {
		return n32
	}

	n64, ok := number.(float64)
	if ok {
		return float32(n64)
	}

	panic("Expected float32 or float64")
}

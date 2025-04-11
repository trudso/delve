package nodes

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Transform struct {
	Position rl.Vector2
	Scale    rl.Vector2
	Rotation rl.Vector2
	
}

func NewTransform() Transform {
	return Transform{
		Position: rl.Vector2{},
		Scale:    rl.Vector2{X: 1, Y: 1},
		Rotation: rl.Vector2{},
	}
}


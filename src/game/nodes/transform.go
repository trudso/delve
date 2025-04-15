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

func (t *Transform) ApplyDataSet(data map[string]any) {
	if d, found := data["position.x"]; found {
		t.Position.X = d.(float32)
	}

	if d, found := data["position.y"]; found {
		t.Position.Y = d.(float32)
	}

	if d, found := data["scale.x"]; found {
		t.Scale.X = d.(float32)
	}

	if d, found := data["scale.y"]; found {
		t.Scale.Y = d.(float32)
	}

	if d, found := data["rotation.x"]; found {
		t.Rotation.X = d.(float32)
	}

	if d, found := data["rotation.y"]; found {
		t.Rotation.Y = d.(float32)
	}
}

func TransformToDataSet(transform Transform, changedFieldsOnly bool) map[string]any {
	res := map[string]any{
		"position.x": transform.Position.X,
		"position.y": transform.Position.Y,
		"scale.x":    transform.Scale.X,
		"scale.y":    transform.Scale.Y,
		"rotation.x": transform.Rotation.X,
		"rotation.y": transform.Rotation.Y,
	}
	return res
}

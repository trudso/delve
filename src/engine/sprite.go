package engine

import (
	"reflect"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Sprite struct {
	BaseNode
	Source  string
	Texture rl.Texture2D
}

func (s Sprite) Draw() {
	rl.DrawTexture(s.Texture, 0, 0, rl.White)
}

func (s Sprite) Delete() {
	rl.UnloadTexture(s.Texture)
	s.BaseNode.Delete()
}

func NewSprite(id string, source string) Sprite {
	sprite := Sprite{
		BaseNode: NewBaseNode(id, reflect.TypeOf(Sprite{}), newSpriteFromDataSet),
		Source:   source,
		Texture:  rl.LoadTexture(source),
	}

	sprite.Transform.Position = rl.Vector2{X: -float32(sprite.Texture.Width / 2), Y: -float32(sprite.Texture.Height / 2)}
	return sprite
}

func newSpriteFromDataSet(id string, data map[string]any) Node {
	source, _ := data["source"].(string)
	result := NewSprite(id, source)
	return &result
}

package nodes

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const SPRITE_NODE = "Sprite"

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
		BaseNode: NewBaseNode(SPRITE_NODE, id),
		Source:   source,
		Texture:  rl.LoadTexture(source),
	}

	sprite.Transform.Position = rl.Vector2{X: -float32(sprite.Texture.Width / 2), Y: -float32(sprite.Texture.Height / 2)}
	return sprite
}

func NewSpriteFromDataSet(data map[string]any) Node {
	id, found := data["id"]
	if !found {
		panic( "No id specified for sprite")
	}

	source, found := data["source"]
	if !found {
		panic( "No source specified for sprite")
	}

	sprite := NewSprite(id.(string), source.(string))
	sprite.ApplyDataSet(data)

	return &sprite
}

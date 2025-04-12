package nodes

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Sprite struct {
	BaseNode
	Source string
	Texture rl.Texture2D
}

func (s Sprite) Draw() {
	rl.DrawTexture(s.Texture, 0, 0, rl.White)
}

func (s Sprite) Close() {
	rl.UnloadTexture(s.Texture)
}

func NewSprite(source string) Sprite {
	sprite := Sprite{
		BaseNode:   NewBaseNode(source),
		Source: source,
		Texture: rl.LoadTexture(source),
	}

	sprite.Transform.Position = rl.Vector2{X: -float32(sprite.Texture.Width/2), Y: -float32(sprite.Texture.Height/2)}
	return sprite
}

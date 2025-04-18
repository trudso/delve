package scenes

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/trudso/delve/engine"
)

const PLAYER_NODE = "Player"

type Player struct {
	// nodes
	engine.BaseNode
	bodySprite      *engine.Sprite
	leftHandSprite  *engine.Sprite
	rightHandSprite *engine.Sprite

	// attributes
	Speed float32
}

func (p *Player) Move(deltaTime float32) {
	if rl.IsKeyDown(rl.KeyUp) {
		p.Transform.Position.Y -= p.Speed * deltaTime
	}
	if rl.IsKeyDown(rl.KeyDown) {
		p.Transform.Position.Y += p.Speed * deltaTime
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		p.Transform.Position.X -= p.Speed * deltaTime
	}
	if rl.IsKeyDown(rl.KeyRight) {
		p.Transform.Position.X += p.Speed * deltaTime
	}

	p.Transform.Rotation.X += 20 * deltaTime
}

func (p Player) GetDataSet(onlyChangedFields bool) map[string]any {
	result := p.BaseNode.GetDataSet(onlyChangedFields)
	result["speed"] = p.Speed
	return result
}


func NewPlayer() Player {
	bodySprite := engine.NewSprite("body", "res/players/green_character.png")
	leftHandSprite := engine.NewSprite("left_hand", "res/players/green_hand.png")
	rightHandSprite := engine.NewSprite("right_hand", "res/players/green_hand.png")
	leftHandSprite.Transform.Position.X += 30
	rightHandSprite.Transform.Position.X -= 30

	player := Player{
		BaseNode:        engine.NewBaseNode(PLAYER_NODE, "player1"),
		bodySprite:      &bodySprite,
		leftHandSprite:  &leftHandSprite,
		rightHandSprite: &rightHandSprite,

		Speed: 100,
	}

	player.AddChild(player.bodySprite)
	player.AddChild(player.leftHandSprite)
	player.AddChild(player.rightHandSprite)

	return player
}

func NewPlayerFromDataSet( data map[string]any) engine.Node {
	player := NewPlayer()
	player.ApplyDataSet(data)
	return &player
}

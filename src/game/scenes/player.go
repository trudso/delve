package scenes

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/trudso/delve/game/nodes"
)

const PLAYER_NODE = "player"

type Player struct {
	// nodes
	nodes.BaseNode
	bodySprite      *nodes.Sprite
	leftHandSprite  *nodes.Sprite
	rightHandSprite *nodes.Sprite

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
	bodySprite := nodes.NewSprite("body", "res/players/green_character.png")
	leftHandSprite := nodes.NewSprite("left_hand", "res/players/green_hand.png")
	rightHandSprite := nodes.NewSprite("right_hand", "res/players/green_hand.png")
	leftHandSprite.Transform.Position.X += 30
	rightHandSprite.Transform.Position.X -= 30

	player := Player{
		BaseNode:        nodes.NewBaseNode(PLAYER_NODE, "player1"),
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

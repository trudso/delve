package levels

import (
	"github.com/trudso/delve/game/nodes"
	"github.com/trudso/delve/game/scenes"
)

type TestLevel struct {
	nodes.BaseNode
	Player *scenes.Player
}

func NewTestLevel() TestLevel {
	player := scenes.NewPlayer()
	level := TestLevel{
		BaseNode: nodes.NewBaseNode("TestLevel"),
		Player:   &player,
	}

	level.Player.Transform.Position.X = 100
	level.Player.Transform.Position.Y = 100

	level.AddChild(level.Player)

	return level
}

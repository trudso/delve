package levels

import (
	"github.com/trudso/delve/game/nodes"
	"github.com/trudso/delve/game/scenes"
)

type TestLevel struct {
	nodes.BaseNode
	Player   *scenes.Player
	Snapshot *scenes.Snapshot
}

func NewTestLevel() TestLevel {
	player := scenes.NewPlayer()
	snapshot := scenes.NewSnapshot("/")

	level := TestLevel{
		BaseNode: nodes.NewBaseNode("TestLevel"),
		Player:   &player,
		Snapshot: &snapshot,
	}

	level.Player.Transform.Position.X = 100
	level.Player.Transform.Position.Y = 100

	level.AddChild(level.Player)
	level.AddChild(level.Snapshot)

	return level
}

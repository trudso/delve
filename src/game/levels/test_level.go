package levels

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/trudso/delve/game/nodes"
	"github.com/trudso/delve/game/scenes"
)

const TESTLEVEL_NODE = "TestLevel"

type TestLevel struct {
	nodes.BaseNode
	Player   *scenes.Player
	Snapshot *scenes.Snapshot
}

func NewTestLevel() TestLevel {
	player := scenes.NewPlayer()
	snapshot := scenes.NewSnapshot("ss1", "/")

	level := TestLevel{
		BaseNode: nodes.NewBaseNode(TESTLEVEL_NODE, "TestLevel1"),
		Player:   &player,
		Snapshot: &snapshot,
	}

	level.Player.Transform.Position.X = 100
	level.Player.Transform.Position.Y = 100

	level.AddChild(level.Player)
	level.AddChild(level.Snapshot)

	return level
}

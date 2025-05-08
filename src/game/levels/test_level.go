package levels

import (
	"github.com/trudso/delve/engine"
	"github.com/trudso/delve/game/scenes"
)

const TESTLEVEL_NODE = "TestLevel"

type TestLevel struct {
	engine.BaseNode
	Player   *scenes.Player
	Snapshot *engine.Snapshot
}

func NewTestLevel() TestLevel {
	player := scenes.NewPlayer()
	snapshot := engine.NewSnapshot("ss1", "/")

	level := TestLevel{
		BaseNode: engine.NewBaseNode(TESTLEVEL_NODE, "TestLevel1"),
		Player:   &player,
		Snapshot: &snapshot,
	}

	level.Player.Transform.Position.X = 100
	level.Player.Transform.Position.Y = 100

	level.AddChild(level.Player)
	level.AddChild(level.Snapshot)

	return level
}

//func NewTestLevelFromDataSet( data map[string]any) engine.Node {
//	level := NewTestLevel()
//	level.ApplyDataSet(data)
//	return &level
//}

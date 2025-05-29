package levels

import (
	"reflect"

	"github.com/trudso/delve/engine"
	"github.com/trudso/delve/game/scenes"
)

type TestLevel struct {
	engine.BaseNode

	Player   *scenes.Player
	Snapshot *engine.Snapshot
}

func NewTestLevel(id string) TestLevel {
	player := scenes.NewPlayer("player")
	snapshot := engine.NewSnapshot("ss1", "/")

	level := TestLevel{
		BaseNode: engine.NewBaseNode(id, reflect.TypeOf(TestLevel{}), newTestLevelFromDataSet),
		Player:   &player,
		Snapshot: &snapshot,
	}

	level.Player.Transform.Position.X = 100
	level.Player.Transform.Position.Y = 100

	level.AddChild(level.Player)
	level.AddChild(level.Snapshot)

	return level
}

func newTestLevelFromDataSet(id string, _ map[string]any) engine.Node {
	node := NewTestLevel(id)
	return &node
}

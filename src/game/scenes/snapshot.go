package scenes

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/trudso/delve/game/nodes"
)

type Snapshot struct {
	nodes.BaseNode

	rootNodeName string
	rootNode     nodes.Node // the root node to replicate
}

func NewSnapshot(rootNodeName string) Snapshot {
	return Snapshot{
		rootNodeName: rootNodeName,
	}
}

func (s *Snapshot) GetRootNode() nodes.Node {
	if s.rootNode == nil {
		s.rootNode = nodes.GetGameContext().GetNodeTree().GetNode(s.rootNodeName)
	}

	return s.rootNode
}

func (s *Snapshot) Input() {
	if rl.IsKeyDown(rl.KeyLeftShift) {
		if rl.IsKeyReleased(rl.KeyOne) {
			s.SaveSnapshot("snapshot1")
		}
	}
}

func (s Snapshot) SaveSnapshot(name string) {
	data := nodes.NodeToDataSet(s.GetRootNode(), false)
	rl.TraceLog(rl.LogDebug, "snapshot %s: %+v", name, data)
}

func LoadSnapshot(name string) nodes.Node {
	rl.TraceLog(rl.LogDebug, "loading snapshot %s", name)
	return nil
}

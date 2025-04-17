package scenes

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/trudso/delve/game/comm"
	"github.com/trudso/delve/game/nodes"
)

const SNAPSHOT_NODE = "Snapshot"

type Snapshot struct {
	nodes.BaseNode

	rootDirectory string
	rootNodeName  string
	RootNode      nodes.Node // the root node to replicate
}

func NewSnapshot(id string, rootNodeName string) Snapshot {
	return Snapshot{
		BaseNode:      nodes.NewBaseNode(SNAPSHOT_NODE, id),
		rootNodeName:  rootNodeName,
		rootDirectory: "snapshots",
	}
}

func (s *Snapshot) Init() {
	s.RootNode = nodes.GetGameContext().GetNodeTree().GetNode(s.rootNodeName)
}

func (s *Snapshot) Input() {
	if rl.IsKeyDown(rl.KeyLeftShift) {
		if rl.IsKeyReleased(rl.KeyOne) {
			s.SaveSnapshot("snapshot1.data")
		}
	}
}

func (s Snapshot) SaveSnapshot(name string) {
	mapData := nodes.NodeToDataSet(s.RootNode, false)
	comm.SaveAsJson(s.rootDirectory, name, mapData)
}

func LoadSnapshot(name string) nodes.Node {
	rl.TraceLog(rl.LogDebug, "loading snapshot %s", name)
	return nil
}

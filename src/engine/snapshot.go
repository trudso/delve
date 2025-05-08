package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const SNAPSHOT_NODE = "Snapshot"

type Snapshot struct {
	BaseNode

	rootDirectory string
	rootNodeName  string
	RootNode      Node // the root node to replicate
}

func NewSnapshot(id string, rootNodeName string) Snapshot {
	return Snapshot{
		BaseNode:      NewBaseNode(SNAPSHOT_NODE, id),
		rootNodeName:  rootNodeName,
		rootDirectory: "snapshots",
	}
}

func (s *Snapshot) Init() {
	s.RootNode = GetGameContext().GetNodeTree().GetNode(s.rootNodeName)
}

func (s *Snapshot) Input() {
	if rl.IsKeyDown(rl.KeyLeftShift) && rl.IsKeyDown(rl.KeyLeftControl) {
		// save
		if rl.IsKeyReleased(rl.KeyOne) {
			s.SaveSnapshot("snapshot1.data")
		}
		if rl.IsKeyReleased(rl.KeyTwo) {
			s.SaveSnapshot("snapshot2.data")
		}
	}

	if rl.IsKeyDown( rl.KeyLeftShift ) && !rl.IsKeyDown( rl.KeyLeftControl ) {
		// load
		if rl.IsKeyReleased(rl.KeyOne) {
			s.LoadSnapshot("snapshot1.data")
		}
		if rl.IsKeyReleased(rl.KeyTwo) {
			s.LoadSnapshot("snapshot2.data")
		}
	}
}

func (s Snapshot) SaveSnapshot(name string) {
	//mapData := NodeToDataSet(s.RootNode, false)
	//SaveJson(s.rootDirectory, name, mapData)
}

func (s Snapshot) LoadSnapshot(name string) {
	//mapData := LoadJson(s.rootDirectory, name)
	//node := DataSetToNode(mapData)
	//GetGameContext().GetNodeTree().SetRootNode(node)
}

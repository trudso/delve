package engine

import (
	"reflect"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Snapshot struct {
	BaseNode

	rootDirectory string
	rootNodeName  string
	RootNode      Node // the root node to replicate
}

func NewSnapshot(id string, rootNodeName string) Snapshot {
	return Snapshot{
		BaseNode:      NewBaseNode(id, reflect.TypeOf(Snapshot{}), newSnapshotFromDataSet),
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

func newSnapshotFromDataSet(id string, data map[string]any) Node {
	rootNodeName, _ := data["rootNodeName"].(string)
	result := NewSnapshot(id, rootNodeName)
	return &result;
}

package scenes

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/trudso/delve/game/nodes"
)

type Snapshot struct {
	nodes.BaseNode
}

func NewSnapshot() Snapshot {
	return Snapshot{
		nodes.NewBaseNode("snapshot"),
	}
}

func (s *Snapshot) Input() {
	if rl.IsKeyDown(rl.KeyLeftShift) {
		if rl.IsKeyReleased(rl.KeyOne) {
			LoadSnapshot("Snapshot1")
		}
	}
}

func LoadSnapshot(name string) nodes.Node {
	rl.TraceLog( rl.LogDebug, "loading snapshot %s", name)
	return nil
}

package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/trudso/delve/engine"
	"github.com/trudso/delve/game/levels"
)

func main() {
	rl.InitWindow(800, 600, "Delve")
	rl.SetTraceLogLevel(rl.LogDebug)
	defer rl.CloseWindow()

	rl.SetTargetFPS(100)

	// create test level
	level := levels.NewTestLevel("TestLevel1")
	replication := level.GetReplication()
	defer engine.DeleteNode(&level)

	// create game context
	createGameContext(&level)
	engine.InitNode(&level)

	// game loop
	dataChanged := false
	for !rl.WindowShouldClose() {
		replication.ResetToChanged()
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText(fmt.Sprintf("data changed: %+v, player pos: %+v", dataChanged, level.Player.Transform.Position), 100, 20, 20, rl.LightGray)

		engine.Update(engine.GetGameContext().GetNodeTree().GetRootNode(), rl.GetFrameTime())

		rl.EndDrawing()
		ds := replication.BuildChangeSet()
		dataChanged = len(ds) != 0
	}
}

func createGameContext(rootNode engine.Node) {
	nodeCreator := engine.NewBaseNodeCreator()
	nodeTree := engine.NewBaseNodeTree()
	engine.NewGameContext(nodeCreator, &nodeTree)

	engine.GetGameContext().GetNodeTree().SetRootNode(rootNode)
}

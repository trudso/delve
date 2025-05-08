package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/trudso/delve/game/levels"
	"github.com/trudso/delve/engine"
)

func main() {
	rl.InitWindow(800, 600, "Delve")
	rl.SetTraceLogLevel(rl.LogDebug)
	defer rl.CloseWindow()

	rl.SetTargetFPS(100)

	// create test level
	level := levels.NewTestLevel()
	defer engine.DeleteNode(&level)

	// create game context
	createGameContext(&level)

	engine.InitNode(&level)

	// game loop
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText(fmt.Sprintf("player pos: %+v", level.Player.Transform.Position), 100, 20, 20, rl.LightGray)

		engine.Update(engine.GetGameContext().GetNodeTree().GetRootNode(), rl.GetFrameTime())

		rl.EndDrawing()
	}
}

func createGameContext(rootNode engine.Node) {
	nodeCreator := engine.NewBaseNodeCreator()
	//nodeCreator.Register(scenes.PLAYER_NODE, scenes.NewPlayerFromDataSet)
	//nodeCreator.Register(levels.TESTLEVEL_NODE, levels.NewTestLevelFromDataSet)

	//TODO[mt]: Add node instantiators for levels and scenes

	nodeTree := engine.NewBaseNodeTree()
	engine.NewGameContext(nodeCreator, &nodeTree)

	engine.GetGameContext().GetNodeTree().SetRootNode(rootNode)
}

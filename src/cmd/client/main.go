package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/trudso/delve/game/levels"
	"github.com/trudso/delve/game/nodes"
)

func main() {
	rl.InitWindow(800, 600, "Delve")
	rl.SetTraceLogLevel(rl.LogDebug)
	defer rl.CloseWindow()

	rl.SetTargetFPS(100)

	// create test level
	level := levels.NewTestLevel()
	defer nodes.DeleteNode(&level)

	// create game context
	createGameContext(&level)

	nodes.InitNode(&level)

	// game loop
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText(fmt.Sprintf("player pos: %+v", level.Player.Transform.Position), 100, 20, 20, rl.LightGray)

		nodes.Update(&level, rl.GetFrameTime())

		rl.EndDrawing()
	}
}

func createGameContext(rootNode nodes.Node) {
	nodeCreator := nodes.NewBaseNodeCreator()
	//TODO[mt]: Add node instantiators for levels and scenes

	nodeTree := nodes.NewBaseNodeTree()
	nodes.NewGameContext(nodeCreator, &nodeTree)

	nodes.GetGameContext().GetNodeTree().SetRootNode(rootNode)
}

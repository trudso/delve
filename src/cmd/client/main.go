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

	// create game context and level
	createGameContext()
	level := createLevel()
	nodes.GetGameContext().GetNodeTree().SetRootNode(&level)
	defer nodes.Close(&level)

	// game loop
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText(fmt.Sprintf("player pos: %+v", level.Player.Transform.Position), 100, 20, 20, rl.LightGray)

		nodes.Update(&level, rl.GetFrameTime())

		rl.EndDrawing()
	}
}

func createLevel() levels.TestLevel {
	return levels.NewTestLevel()
}

func createGameContext() {
	nodeCreator := nodes.NewBaseNodeCreator()
	//TODO[mt]: Add node instantiators for levels and scenes

	nodeTree := nodes.NewBaseNodeTree()
	nodes.NewGameContext(nodeCreator, &nodeTree)
}

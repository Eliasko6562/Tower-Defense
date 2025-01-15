package main

import rl "github.com/gen2brain/raylib-go/raylib"



func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	game := game{}
	game.init()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		game.update()

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.BeginMode3D(game.camera)
		


		rl.DrawCube(rl.NewVector3(0, 0, 0), 2, 2, 2, rl.Red)
		rl.DrawCubeWires(rl.NewVector3(0, 0, 0), 2, 2, 2, rl.Maroon)
		rl.DrawGrid(10, 1.0)








		rl.EndMode3D()
		rl.EndDrawing()
	}
}
package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	game := game{}
	game.init()

	mesh := rl.GenMeshCube(1.0, 1.0, 1.0)
	model := rl.LoadModelFromMesh(mesh)

	image := rl.GenImageWhiteNoise(128, 128, 16)
	texture := rl.LoadTextureFromImage(image)
	rl.UnloadImage(image)
	model.Materials.Maps.Texture = texture

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		game.update()

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.BeginMode3D(game.camera)

		rl.DrawModel(model, rl.NewVector3(0, 0, 0), 1, rl.Red)
		rl.DrawGrid(10, 1.0)

		rl.EndMode3D()
		rl.EndDrawing()
	}

	rl.UnloadTexture(texture)
}

package main

import rl "github.com/gen2brain/raylib-go/raylib"

type game struct {
	// The camera to use for rendering.
	camera      rl.Camera3D
	cameraSpeed float32
}

func (g *game) init() {
	// Initialize the camera.
	g.camera = rl.Camera3D{}
	g.camera.Position = rl.NewVector3(10.0, 10.0, 10.0)
	g.camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
	g.camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	g.camera.Fovy = 10.0
	g.camera.Projection = rl.CameraOrthographic
	g.cameraSpeed = 1.0
}

func (g *game) update() {
	// Update the camera.
	if rl.IsKeyDown(rl.KeyW) {
		g.camera.Position.X += g.cameraSpeed
		g.camera.Target.X += g.cameraSpeed
		g.camera.Position.Z += g.cameraSpeed
		g.camera.Target.Z += g.cameraSpeed
	}
	if rl.IsKeyDown(rl.KeyS) {
		g.camera.Position.X -= g.cameraSpeed
		g.camera.Target.X -= g.cameraSpeed
		g.camera.Position.Z -= g.cameraSpeed
		g.camera.Target.Z -= g.cameraSpeed
	}
	if rl.IsKeyDown(rl.KeyA) {
		g.camera.Position.X -= g.cameraSpeed
		g.camera.Target.X -= g.cameraSpeed
		g.camera.Position.Z += g.cameraSpeed
		g.camera.Target.Z += g.cameraSpeed
	}
	if rl.IsKeyDown(rl.KeyD) {
		g.camera.Position.X += g.cameraSpeed
		g.camera.Target.X += g.cameraSpeed
		g.camera.Position.Z -= g.cameraSpeed
		g.camera.Target.Z -= g.cameraSpeed
	}
	if rl.GetMouseWheelMove() != 0 {
		g.camera.Position.X += g.cameraSpeed
		g.camera.Target.X += g.cameraSpeed
		g.camera.Position.Y += g.cameraSpeed
		g.camera.Target.Y += g.cameraSpeed
		g.camera.Position.Z += g.cameraSpeed
		g.camera.Target.Z += g.cameraSpeed
		g.camera.Fovy += g.cameraSpeed
	}

}

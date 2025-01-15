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
	g.cameraSpeed = 0.075
}

func (g *game) update() {
	finalSpeed := g.cameraSpeed * g.camera.Fovy
	// Update the camera.
	if rl.IsKeyDown(rl.KeyW) {
		g.camera.Position.X -= finalSpeed
		g.camera.Target.X -= finalSpeed
		g.camera.Position.Z -= finalSpeed
		g.camera.Target.Z -= finalSpeed
	}
	if rl.IsKeyDown(rl.KeyS) {
		g.camera.Position.X += finalSpeed
		g.camera.Target.X += finalSpeed
		g.camera.Position.Z += finalSpeed
		g.camera.Target.Z += finalSpeed
	}
	if rl.IsKeyDown(rl.KeyA) {
		g.camera.Position.X -= finalSpeed
		g.camera.Target.X -= finalSpeed
		g.camera.Position.Z += finalSpeed
		g.camera.Target.Z += finalSpeed
	}
	if rl.IsKeyDown(rl.KeyD) {
		g.camera.Position.X += finalSpeed
		g.camera.Target.X += finalSpeed
		g.camera.Position.Z -= finalSpeed
		g.camera.Target.Z -= finalSpeed
	}
	if scroll := rl.GetMouseWheelMove(); scroll != 0 {
		if scroll > 0 && g.camera.Fovy > 3 {
			g.camera.Position.X -= finalSpeed
			g.camera.Target.X -= finalSpeed
			g.camera.Position.Y -= finalSpeed
			g.camera.Target.Y -= finalSpeed
			g.camera.Position.Z -= finalSpeed
			g.camera.Target.Z -= finalSpeed
			g.camera.Fovy -= finalSpeed
		} else if scroll < 0 && g.camera.Fovy < 100 {
			g.camera.Position.X += finalSpeed
			g.camera.Target.X += finalSpeed
			g.camera.Position.Y += finalSpeed
			g.camera.Target.Y += finalSpeed
			g.camera.Position.Z += finalSpeed
			g.camera.Target.Z += finalSpeed
			g.camera.Fovy += finalSpeed
		}
	}

}

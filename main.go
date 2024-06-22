package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/tinarao/gorl/camera"
	"github.com/tinarao/gorl/music"
	"github.com/tinarao/gorl/player"
	"github.com/tinarao/gorl/sprites"
	"github.com/tinarao/gorl/tilemap"
)

const (
	ScreenWidth  = 1000
	ScreenHeight = 480
)

var (
	running = true
	pl      *player.Player
	cam     *camera.Camera
	tileMap *tilemap.Map
)

func drawScene() {
	tileMap.Draw()
	rl.DrawTexturePro(sprites.PlayerTexture, pl.SrcRect, pl.DstRect, rl.NewVector2(pl.DstRect.Width, pl.DstRect.Height), 0, rl.White)
}

func drawDebugInfo() {
	fpsStr := fmt.Sprintf("FPS: %v", rl.GetFPS())
	playerPosXStr := fmt.Sprintf("player.x: %v", pl.DstRect.X)
	playerPosYStr := fmt.Sprintf("player.y: %v", pl.DstRect.Y)

	isMusicPlStr := fmt.Sprintf("music playing: %v", !music.IsMusicPaused)

	rl.DrawText(fpsStr, 10, 10, 14, rl.Black)
	rl.DrawText(playerPosXStr, 10, 30, 14, rl.Black)
	rl.DrawText(playerPosYStr, 10, 50, 14, rl.Black)
	rl.DrawText(isMusicPlStr, 10, 70, 14, rl.Black)
}

func input() {
	pl.Controls()
	music.Controls()
}

func update() {
	running = !rl.WindowShouldClose()
	music.Play()

	pl.Update()
	cam.FixOnTarget(pl)

	cam.Zoom()
}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.White)
	rl.BeginMode2D(*cam.TwoDim)

	drawScene()
	drawDebugInfo()

	rl.EndMode2D()
	rl.EndDrawing()
}

func gorlInit() {
	rl.InitWindow(ScreenWidth, ScreenHeight, "gorl - cross animaling")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)

	tileMap = tilemap.CreateMap()
	tileMap.InitializeMap()

	music.Init()

	pl = player.Create(sprites.PlayerTexture, 0, 0)
	cam = camera.CreateCamera(ScreenWidth, ScreenHeight, pl)
	sprites.Load()
}

func quit() {
	sprites.Remove()
	music.Quit()

	rl.CloseWindow()
}

func main() {
	gorlInit()

	for running {
		input()
		update()
		render()
	}

	quit()
}

package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
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
	running    = true
	pl         *player.Player
	cam        rl.Camera2D
	frameCount int
	tileMap    *tilemap.Map
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
	if rl.IsKeyPressed(rl.KeyM) {
		music.IsMusicPaused = !music.IsMusicPaused
	}
}

func update() {
	running = !rl.WindowShouldClose()
	//rl.UpdateMusicStream(music.Soundtrack)
	//if music.IsMusicPaused {
	//	rl.PauseMusicStream(music.Soundtrack)
	//} else {
	//	rl.ResumeMusicStream(music.Soundtrack)
	//}

	pl.SrcRect.X = 0
	pl.UpdatePosition()
	if pl.IsMoving && frameCount%8 == 1 {
		pl.Frame++
	}

	frameCount++
	pl.UpdateSprites()

	cam.Target = rl.NewVector2(pl.DstRect.X-(pl.DstRect.Width/2), pl.DstRect.Y-(pl.DstRect.Height/2))

	pl.IsMoving = false
	pl.Up, pl.Down, pl.Right, pl.Left = false, false, false, false

	// Zoom on mouse wheel move
	wheelPos := rl.GetMouseWheelMove() // up=1; down=-1
	if wheelPos > 0 && cam.Zoom < 5 {
		cam.Zoom += 0.5
	} else if wheelPos < 0 && cam.Zoom > 1 {
		cam.Zoom -= 0.5
	}
}

func render() {
	rl.BeginDrawing()
	//rl.ClearBackground(colors.Background())
	rl.ClearBackground(rl.White)
	rl.BeginMode2D(cam)

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
	cam = rl.NewCamera2D(
		rl.NewVector2(float32(ScreenWidth/2), float32(ScreenHeight/2)),
		rl.NewVector2(pl.DstRect.X-(pl.DstRect.Width/2), pl.DstRect.Y-(pl.DstRect.Height/2)),
		0.0,
		1.0)
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

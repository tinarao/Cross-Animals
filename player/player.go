package player

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/tinarao/gorl/helpers"
)

type Player struct {
	Texture rl.Texture2D
	SrcRect rl.Rectangle
	DstRect rl.Rectangle

	Speed     float32
	IsMoving  bool
	Direction int // 1 for left, 2 for right, 3 for up, 4 for down

	Up, Down, Right, Left bool

	Frame int
}

var frameCount = &helpers.FrameCount

func Create(tex rl.Texture2D, x float32, y float32) *Player {
	return &Player{
		Texture:  tex,
		SrcRect:  rl.NewRectangle(0, 0, 48, 48),
		DstRect:  rl.NewRectangle(x, y, 100, 100),
		Speed:    3,
		IsMoving: false,
	}
}

func (pl *Player) Controls() {
	if rl.IsKeyDown(rl.KeyW) {
		pl.IsMoving = true
		pl.Direction = 1
		pl.Up = true
	}
	if rl.IsKeyDown(rl.KeyS) {
		pl.IsMoving = true
		pl.Direction = 0
		pl.Down = true
	}
	if rl.IsKeyDown(rl.KeyD) {
		pl.IsMoving = true
		pl.Direction = 3
		pl.Right = true
	}
	if rl.IsKeyDown(rl.KeyA) {
		pl.IsMoving = true
		pl.Direction = 2
		pl.Left = true
	}
}

func (pl *Player) updateSprites() {
	pl.SrcRect.X = 0
	if pl.Frame > 3 {
		pl.Frame = 0
	}

	if pl.IsMoving {
		pl.SrcRect.X = pl.SrcRect.Width * float32(pl.Frame)
		pl.SrcRect.Y = pl.SrcRect.Height * float32(pl.Direction)
	}
}

func (pl *Player) updatePosition() {
	if pl.IsMoving {
		if pl.Up {
			pl.DstRect.Y -= pl.Speed
		}
		if pl.Down {
			pl.DstRect.Y += pl.Speed
		}
		if pl.Left {
			pl.DstRect.X -= pl.Speed
		}
		if pl.Right {
			pl.DstRect.X += pl.Speed
		}
	}

	if pl.IsMoving && *frameCount%8 == 1 {
		pl.Frame++
	}
}

func (pl *Player) Update() {
	pl.updatePosition()
	*frameCount++
	pl.updateSprites()

	pl.IsMoving = false
	pl.Up, pl.Down, pl.Right, pl.Left = false, false, false, false
}

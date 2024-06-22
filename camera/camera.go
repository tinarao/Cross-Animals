package camera

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/tinarao/gorl/player"
)

type Camera struct {
	TwoDim *rl.Camera2D
}

func CreateCamera(screenw int, screenh int, pl *player.Player) *Camera {
	cam := rl.NewCamera2D(
		rl.NewVector2(float32(screenw/2), float32(screenh/2)),
		rl.NewVector2(pl.DstRect.X-(pl.DstRect.Width/2), pl.DstRect.Y-(pl.DstRect.Height/2)),
		0.0,
		1.0)

	return &Camera{
		TwoDim: &cam,
	}
}

func (cam *Camera) Zoom() {
	wheelPos := rl.GetMouseWheelMove() // up=1; down=-1
	if wheelPos > 0 && cam.TwoDim.Zoom < 5 {
		cam.TwoDim.Zoom += 0.5
	} else if wheelPos < 0 && cam.TwoDim.Zoom > 1 {
		cam.TwoDim.Zoom -= 0.5
	}
}

func (cam *Camera) FixOnTarget(pl *player.Player) {
	cam.TwoDim.Target = rl.NewVector2(pl.DstRect.X-(pl.DstRect.Width/2), pl.DstRect.Y-(pl.DstRect.Height/2))
}

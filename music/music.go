package music

import rl "github.com/gen2brain/raylib-go/raylib"

var IsMusicPaused bool
var Soundtrack rl.Music

func Init() {
	rl.InitAudioDevice()
	Soundtrack = rl.LoadMusicStream("assets/Music/Power.mp3")
	IsMusicPaused = false

	rl.PlayMusicStream(Soundtrack)
}

func Quit() {
	rl.UnloadMusicStream(Soundtrack)
	rl.CloseAudioDevice()
}

func Play() {
	rl.UpdateMusicStream(Soundtrack)
	if IsMusicPaused {
		rl.PauseMusicStream(Soundtrack)
	} else {
		rl.ResumeMusicStream(Soundtrack)
	}
}

func Controls() {
	if rl.IsKeyPressed(rl.KeyM) {
		IsMusicPaused = !IsMusicPaused
	}
}

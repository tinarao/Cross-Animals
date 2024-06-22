package sprites

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	GrassTexture  rl.Texture2D
	PlayerTexture rl.Texture2D
	WaterTexture  rl.Texture2D
)

var LoadedSprites = make([]rl.Texture2D, 2)

func Load() {
	GrassTexture = rl.LoadTexture("assets/Tilesets/Grass.png")
	PlayerTexture = rl.LoadTexture("assets/Characters/BasicCharakterSpritesheet.png")
	WaterTexture = rl.LoadTexture("assets/Tilesets/Water.png")

	_ = append(LoadedSprites, GrassTexture)
	_ = append(LoadedSprites, PlayerTexture)
	_ = append(LoadedSprites, WaterTexture)
}

func Remove() {
	for _, sprite := range LoadedSprites {
		rl.UnloadTexture(sprite)
	}
}

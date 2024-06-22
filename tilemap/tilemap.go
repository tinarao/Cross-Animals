package tilemap

import (
	"encoding/csv"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/tinarao/gorl/sprites"
	"log"
	"os"
)

type Map struct {
	Water  [][]string
	Ground [][]string
}

func CreateMap() *Map {
	return &Map{}
}

func (m *Map) InitializeMap() {
	csvToLoad := [2]string{"assets/Tilemap/map_water.csv", "assets/Tilemap/map_ground.csv"}

	for k, v := range csvToLoad {
		file, err := os.Open(v)
		if err != nil {
			log.Fatal(err)
		}

		r := csv.NewReader(file)

		tiles, err := r.ReadAll()
		if err != nil {
			log.Fatal(err)
		}

		if k == 0 {
			m.Water = tiles
		} else if k == 1 {
			m.Ground = tiles
		}

		file.Close()
	}

}

func (m *Map) Draw() {
	//for x, row := range m.Water {
	//	for y, _ := range row {
	//		rl.DrawTexturePro(
	//			sprites.WaterTexture,
	//			rl.NewRectangle(0, 0, 16, 16),                          // first tile of tileset
	//			rl.NewRectangle(float32(-x*16), float32(y*16), 16, 16), // destination. draws initially at 100, 100 with size 100, 100
	//			rl.NewVector2(0, 0),
	//			0,
	//			rl.White,
	//		)
	//	}
	//}

	for x, row := range m.Ground {
		for y, tile := range row {
			if tile == "12" {
				rl.DrawTexturePro(
					sprites.WaterTexture,
					rl.NewRectangle(0, 0, 16, 16),                          // first tile of tileset
					rl.NewRectangle(float32(-x*16), float32(y*16), 16, 16), // destination. draws initially at 100, 100 with size 100, 100
					rl.NewVector2(float32(x), float32(y)),
					0,
					rl.White,
				)
			}
		}
	}
}

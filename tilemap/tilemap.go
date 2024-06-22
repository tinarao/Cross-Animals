package tilemap

import (
	"encoding/csv"
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/tinarao/gorl/sprites"
	"log"
	"os"
)

type Map struct {
	Water  [][]string
	Ground [][]string

	CreatedTiles []Tile
}

type Tile struct {
	Src     rl.Rectangle
	Dst     rl.Rectangle
	Texture rl.Texture2D
}

const TILESIZE = 100

func CreateMap() *Map {
	return &Map{}
}

// InitializeMap parses the .csv files, then generates a Tile object
// and push it to the Map.CreatedTiles slice
func (m *Map) InitializeMap() {
	csvToLoad := [2]string{"assets/Tilemap/map_water.csv", "assets/Tilemap/map_ground.csv"}

	// Parse csvToLoad
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

		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}

	m.CreatedTiles = make([]Tile, 100)
	// Create Tile instances and push to m.CreatedTiles
	index := 0
	for i, row := range m.Ground {
		for y, val := range row {
			if val != "-1" {
				fmt.Printf("val: %s;\tx: %d;\ty: %d\n", val, i*16, y*16)
				tile := Tile{
					Texture: sprites.GrassTexture,
					Src:     rl.NewRectangle(0, 0, 16, 16),
					Dst:     rl.NewRectangle(float32(i*16), float32(y*16), TILESIZE, TILESIZE),
				}
				m.CreatedTiles[index] = tile
				fmt.Printf("index: %d\n", index)
				index++
			}

			continue
		}
	}

	return
}
func (m *Map) Draw() {
	for _, tile := range m.CreatedTiles {
		rl.DrawTexturePro(
			tile.Texture,
			tile.Src,
			tile.Dst,
			rl.NewVector2(0, 0),
			0,
			rl.White,
		)
	}
}

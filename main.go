package main

import (
	"fmt"

	"github.com/golang/freetype/truetype"
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/imdraw"
	"github.com/gopxl/pixel/v2/pixelgl"
	"github.com/gopxl/pixel/v2/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/gofont/goregular"
)

// Calculates the squares that need to be drawn on the grid
// and pushed them to IMDRaw
// TOOD: Calculate dynamically based on window size
func calculateGridSquares(imd *imdraw.IMDraw) {
	imd.Color = colornames.Lightgray

	offset := 0

	for j := 0; j < 10; j++ {
		for i := 0; i < 5; i++ {
			imd.Push(
				pixel.V(
					float64(100*i+offset),
					float64(j*50),
				),
				pixel.V(
					float64(100*i+50+offset),
					float64(50*j+50),
				),
			)
			imd.Rectangle(0)
		}
		if offset == 0 {
			offset = 50
		} else {
			offset = 0
		}
	}
	imd.Color = colornames.Black
	imd.Push(pixel.V(0, 500), pixel.V(500, 530))
	imd.Rectangle(0)
}

func loadAssets() map[string]*pixel.Sprite {
	assets := make(map[string]*pixel.Sprite)

	pic, err := loadPicture("assets/gopher50x50.png")
	if err != nil {
		panic(err)
	}
	assets["gopher"] = pixel.NewSprite(pic, pic.Bounds())

	pic, err = loadPicture("assets/carrot.png")
	if err != nil {
		panic(err)
	}
	assets["reward"] = pixel.NewSprite(pic, pic.Bounds())

	pic, err = loadPicture("assets/rock.png")
	if err != nil {
		panic(err)
	}
	assets["obstacle"] = pixel.NewSprite(pic, pic.Bounds())

	pic, err = loadPicture("assets/dollar-bag2.png")
	if err != nil {
		panic(err)
	}
	assets["win"] = pixel.NewSprite(pic, pic.Bounds())

	return assets
}

func run() {
	// Create window config
	// TODO: Customizable bounds using cmd args
	config := pixelgl.WindowConfig{
		Title:  "Grid Game",
		Bounds: pixel.R(0, 0, 500, 530),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(config)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)
	calculateGridSquares(imd)

	gopher := NewPlayer(Coord{0, 0})
	grid := GenerateGrid(10, 10, &gopher)

	ttf, err := truetype.Parse(goregular.TTF)
	if err != nil {
		panic(err)
	}

	face := truetype.NewFace(ttf, &truetype.Options{Size: 22})
	atlas := text.NewAtlas(face, text.ASCII)
	txt := text.New(pixel.V(5, 505), atlas)
	txt.Color = colornames.Yellow

	for !win.Closed() && !win.JustPressed(pixelgl.KeyEscape) {
		// Make background white
		win.Clear(colornames.White)

		// Draw the grey squares
		imd.Draw(win)

		// Write the score text to the top of the window
		txt.WriteString(fmt.Sprintf("Score: %.1f", gopher.score))
		txt.Draw(win, pixel.IM)
		txt.Clear()

		// Draw the sprites contained in grid.tiles
		grid.Draw(win)

		// Handle key inputs
		if win.JustPressed(pixelgl.KeyLeft) {
			if grid.IsValidTile(gopher.location.Shift(East)) {
				gopher.Move(East, &grid)
			}
			gopher.direction = -1
		}
		if win.JustPressed(pixelgl.KeyRight) {
			if grid.IsValidTile(gopher.location.Shift(West)) {
				gopher.Move(West, &grid)
			}
			gopher.direction = 1
		}
		if win.JustPressed(pixelgl.KeyUp) {
			if grid.IsValidTile(gopher.location.Shift(North)) {
				gopher.Move(North, &grid)
			}
		}
		if win.JustPressed(pixelgl.KeyDown) {
			if grid.IsValidTile(gopher.location.Shift(South)) {
				gopher.Move(South, &grid)
			}
		}
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}

package main

import (
	"fmt"
	"time"

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

	for !win.Closed() && !win.JustPressed(pixelgl.KeyEscape) && !grid.gameOver {
		// Handle key inputs
		if win.JustPressed(pixelgl.KeyLeft) {
			if grid.IsValidTile(gopher.location.Shift(East)) {
				gopher.Move(East, &grid)
			}
			gopher.direction = East
		}
		if win.JustPressed(pixelgl.KeyRight) {
			if grid.IsValidTile(gopher.location.Shift(West)) {
				gopher.Move(West, &grid)
			}
			gopher.direction = West
		}
		if win.JustPressed(pixelgl.KeyUp) {
			if grid.IsValidTile(gopher.location.Shift(North)) {
				gopher.Move(North, &grid)
			}
			gopher.direction = North
		}
		if win.JustPressed(pixelgl.KeyDown) {
			if grid.IsValidTile(gopher.location.Shift(South)) {
				gopher.Move(South, &grid)
			}
			gopher.direction = South
		}

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

		win.Update()
	}

	txt = text.New(pixel.V(200, 505), atlas)
	txt.Color = colornames.Yellow

	fmt.Printf("You ended the game with a score of: %.1f\n", gopher.score)

	for !win.Closed() && !win.JustPressed(pixelgl.KeyEscape) {
		txt.WriteString("GAME OVER!")
		txt.Draw(win, pixel.IM)
		txt.Clear()
		win.Update()
	}
}

func run_Q() {
	// Create window config
	config := pixelgl.WindowConfig{
		Title:  "Grid Game - Q-Learning Agent",
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

	ttf, _ := truetype.Parse(goregular.TTF)
	face := truetype.NewFace(ttf, &truetype.Options{Size: 22})
	atlas := text.NewAtlas(face, text.ASCII)
	txt := text.New(pixel.V(5, 505), atlas)
	txt.Color = colornames.Yellow

	agent := NewQLearningAgent(4, 0.1, 0.95, 0.5, 0.99)

	win.Clear(colornames.White)
	imd.Draw(win)
	txt.WriteString(fmt.Sprintf("Score: %.1f", gopher.score))
	txt.Draw(win, pixel.IM)
	txt.Clear()
	grid.Draw(win)

	iteration := 0

	for !win.Closed() && !win.JustPressed(pixelgl.KeyEscape) {
		fmt.Printf("Iteration %d\n", iteration)
		// iteration++

		currentGameState := GameState{
			grid: &grid,
			Coord: Coord{
				tileX: gopher.location.tileX,
				tileY: gopher.location.tileY,
			},
		}

		direction := agent.act(currentGameState)
		oldScore := gopher.score
		gopher.Move(direction, &grid)
		gopher.direction = direction

		agent.update(currentGameState, direction, GameState{grid: &grid, Coord: gopher.location}, gopher.score - oldScore)


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

		if iteration > 500 {
			agent.ExplorationRate = 0.0
			time.Sleep(100 * time.Millisecond)
			win.Update()
		}

		if grid.gameOver {
			grid.ResetGrid(&gopher)
			gopher.windowX = 25
			gopher.windowY = 25
			gopher.location = Coord{0, 0}
			gopher.score = 0
			iteration++
		}
	}
}

func main() {
	pixelgl.Run(run_Q)
	// pixelgl.Run(run)
}

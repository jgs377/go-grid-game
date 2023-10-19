package main

import (
	"fmt"
	"image"
	"os"

	_ "image/png"

	"github.com/golang/freetype/truetype"
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/imdraw"
	"github.com/gopxl/pixel/v2/pixelgl"
	"github.com/gopxl/pixel/v2/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/gofont/goregular"
)

type coord struct {
	tileX int
	tileY int
}

type gopher struct {
	windowX   float64
	windowY   float64
	location  coord
	direction float64 // TODO: improve; -1 = left, 1 = right

	score float64
}

type grid struct {
	sizeX     int
	sizeY     int
	rewards   []coord
	obstacles []coord
	finish    []coord
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func generateGrid(imd *imdraw.IMDraw) {
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

func initGrid() grid {
	g := grid{
		sizeX: 10,
		sizeY: 10,
	}
	g.rewards = append(g.rewards, coord{tileX: 5, tileY: 5})
	g.rewards = append(g.rewards, coord{tileX: 3, tileY: 3})
	return g
}

func isObstacle(tile coord, g grid) bool {
	// TODO
	return false
}

func isReward(tile coord, g grid) bool {
	// TODO
	return false
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
	cfg := pixelgl.WindowConfig{
		Title:  "Grid Game",
		Bounds: pixel.R(0, 0, 500, 530),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.Clear(colornames.White)
	win.SetSmooth(true)

	imd := imdraw.New(nil)
	generateGrid(imd)

	// pic, err := loadPicture("assets/gopher50x50.png")
	// if err != nil {panic(err)}
	// gopherSprite := pixel.NewSprite(pic, pic.Bounds())

	assets := loadAssets()

	// state := gopher{26, 27, coord{0, 0}}
	state := gopher{25, 25, coord{0, 0}, -1.0, 0}
	toggle := true

	ttf, err := truetype.Parse(goregular.TTF)
	if err != nil {
		panic(err)
	}
	face := truetype.NewFace(ttf, &truetype.Options{
		Size: 22,
	})
	atlas := text.NewAtlas(face, text.ASCII)

	txt := text.New(pixel.V(5, 505), atlas)
	txt.Color = colornames.Yellow

	for !win.Closed() && !win.JustPressed(pixelgl.KeyEscape) {

		// Empty white canvas
		win.Clear(colornames.White)

		// Draw black squares
		imd.Draw(win)

		txt.WriteString(fmt.Sprintf("Score: %.1f", state.score))
		txt.Draw(win, pixel.IM)
		txt.Clear()

		// Draw gopher
		mat := pixel.IM
		// mat = mat.Scaled(pixel.ZV, 0.08)
		mat = mat.ScaledXY(pixel.ZV, pixel.V(state.direction, 1))
		mat = mat.Moved(pixel.V(state.windowX, state.windowY))
		assets["gopher"].Draw(win, mat)

		if win.JustPressed(pixelgl.KeyLeft) {
			if state.location.tileX > 0 {
				state.windowX -= 50
				state.location.tileX -= 1
				state.score -= 0.1
			}
			state.direction = -1
		}
		if win.JustPressed(pixelgl.KeyRight) {
			if state.location.tileX < 9 {
				state.windowX += 50
				state.location.tileX += 1
				state.score -= 0.1
			}
			state.direction = 1
		}
		if win.JustPressed(pixelgl.KeyUp) {
			if state.location.tileY < 9 {
				state.windowY += 50
				state.location.tileY += 1
				state.score -= 0.1
			}
		}
		if win.JustPressed(pixelgl.KeyDown) {
			if state.location.tileY > 0 {
				state.windowY -= 50
				state.location.tileY -= 1
				state.score -= 0.1
			}
		}

		// Update canvas
		if toggle {
			win.Update()
			toggle = false
		} else {
			win.SwapBuffers()
			win.UpdateInputWait(0)
		}
	}
}

func main() {
	pixelgl.Run(run)
}

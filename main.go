package gogridgame

import (
	"image"
	"os"

	_ "image/png"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/imdraw"
	"github.com/gopxl/pixel/v2/pixelgl"
	"golang.org/x/image/colornames"
)

type gopher struct {
	X     float64
	Y     float64
	tileX int
	tileY int
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

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Grid Game",
		Bounds: pixel.R(0, 0, 500, 500),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.Clear(colornames.White)
	win.SetSmooth(true)

	imd := imdraw.New(nil)
	imd.Color = colornames.Black
	// imd.Push(pixel.V(0, 0), pixel.V(50, 50))
	// imd.Rectangle(0)

	// toggle := true

	x := 0
	offset := 0

	for j := 0; j < 10; j++ {
		for i := 0; i < 5; i++ {
			// if toggle{
			// 	imd.Color = colornames.Black
			// } else {
			// 	imd.Color = colornames.White
			// }
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

	pic, err := loadPicture("hiking.png")
	if err != nil {
		panic(err)
	}
	sprite := pixel.NewSprite(pic, pic.Bounds())
	// angle := 0.0

	state := gopher{26, 27, 0, 0}
	// movement := true

	toggle := true

	for !win.Closed() && !win.JustPressed(pixelgl.KeyEscape) {
		// Empty white canvas
		win.Clear(colornames.White)

		// Draw black squares
		imd.Draw(win)

		// Draw gopher
		mat := pixel.IM
		mat = mat.Scaled(pixel.ZV, 0.08)
		mat = mat.Moved(pixel.V(state.X, state.Y))
		sprite.Draw(win, mat)

		if win.JustPressed(pixelgl.KeyLeft) {
			if state.tileX > 0 {
				state.X -= 50
				state.tileX -= 1
			}
		}
		if win.JustPressed(pixelgl.KeyRight) {
			if state.tileX < 9 {
				state.X += 50
				state.tileX += 1
			}
		}
		if win.JustPressed(pixelgl.KeyUp) {
			if state.tileY < 9 {
				state.Y += 50
				state.tileY += 1
			}
		}
		if win.JustPressed(pixelgl.KeyDown) {
			if state.tileY > 0 {
				state.Y -= 50
				state.tileY -= 1
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

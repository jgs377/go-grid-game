package main

import (
	"fmt"
	"image"
	"os"

	_ "image/png"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/pixelgl"
)

const debug bool = false

type object struct {
	// Generic object in the grid game
	// Can be a gopher, obstacle, reward,
	windowX  float64
	windowY  float64
	location Coord

	sprite *pixel.Sprite
}

type Player struct {
	object

	score     float64
	direction int

	backSprite  *pixel.Sprite
	frontSprite *pixel.Sprite
}

type Obstacle struct {
	object

	// In case we allow breaking obstacles in the future
	// breakable    bool
	// broken       bool
	// breakPenalty float64
	// walkPenalty  float64
}

type Reward struct {
	object
	value int
}

type LossCondition struct {
	object
	value int
}

type WinCondition struct {
	object
	value int
}

// All object types present in the game implement this ObjectInterface
type ObjectInterface interface {
	Draw(*pixelgl.Window)
}

func (p Player) Draw(win *pixelgl.Window) {
	// Check East first because it's the only one we need to scale
	// the matrix for which has to be step 1
	if p.direction == East {
		mat := pixel.IM.ScaledXY(pixel.ZV, pixel.V(-1, 1))
		mat = mat.Moved(pixel.V(p.windowX, p.windowY))
		p.sprite.Draw(win, mat)
		return
	}

	mat := pixel.IM.Moved(pixel.V(p.windowX, p.windowY))
	if p.direction == North {
		p.backSprite.Draw(win, mat)
		return
	} else if p.direction == South {
		p.frontSprite.Draw(win, mat)
		return
	}
	
	p.sprite.Draw(win, mat)
}

func (o Obstacle) Draw(win *pixelgl.Window) {
	mat := pixel.IM.Moved(pixel.V(o.windowX, o.windowY))
	o.sprite.Draw(win, mat)
}

func (r Reward) Draw(win *pixelgl.Window) {
	// TODO
}

func (w WinCondition) Draw(win *pixelgl.Window) {
	// TODO
}

func (l LossCondition) Draw(win *pixelgl.Window) {
	// TODO
}

// Moves the player in the grid
func (p *Player) Move(direction int, grid *Grid) {
	if direction == North {
		grid.tiles[p.location.tileX][p.location.tileY+1] = grid.tiles[p.location.tileX][p.location.tileY]
		grid.tiles[p.location.tileX][p.location.tileY] = nil
		p.location.tileY += 1
		p.windowY += 50
		p.score -= 0.1
	}
	if direction == South {
		grid.tiles[p.location.tileX][p.location.tileY-1] = grid.tiles[p.location.tileX][p.location.tileY]
		grid.tiles[p.location.tileX][p.location.tileY] = nil
		p.location.tileY -= 1
		p.windowY -= 50
		p.score -= 0.1
	}
	if direction == East {
		grid.tiles[p.location.tileX-1][p.location.tileY] = grid.tiles[p.location.tileX][p.location.tileY]
		grid.tiles[p.location.tileX][p.location.tileY] = nil
		p.location.tileX -= 1
		p.windowX -= 50
		p.score -= 0.1
	}
	if direction == West {
		grid.tiles[p.location.tileX+1][p.location.tileY] = grid.tiles[p.location.tileX][p.location.tileY]
		grid.tiles[p.location.tileX][p.location.tileY] = nil
		p.location.tileX += 1
		p.windowX += 50
		p.score -= 0.1
	}
	if debug {
		for _, j := range grid.tiles {
			fmt.Println(j)
		}
		fmt.Println("================")
	}
}

// Generates an object of type Player
func NewPlayer(coord Coord) (player Player) {
	pic, err := loadPicture("assets/gopher50x50.png")
	if err != nil {
		panic(err)
	}
	pic2, err := loadPicture("assets/gopher50x50back.png")
	if err != nil {
		panic(err)
	}
	pic3, err := loadPicture("assets/gopher50x50front.png")
	if err != nil {
		panic(err)
	}

	player = Player{
		object: object{
			windowX:  25,
			windowY:  25,
			location: coord,
			sprite:   pixel.NewSprite(pic, pic.Bounds()),
		},
		score:       0,
		direction:   1.0,
		backSprite:  pixel.NewSprite(pic2, pic2.Bounds()),
		frontSprite: pixel.NewSprite(pic3, pic3.Bounds()),
	}

	return player
}

// Generates an object of type Obstacle
func NewObstacle(coord Coord) (obstacle Obstacle) {
	pic, err := loadPicture("assets/rock.png")
	if err != nil {
		panic(err)
	}

	obstacle = Obstacle{
		object: object{
			windowX:  float64(25 + 50*coord.tileX),
			windowY:  float64(25 + 50*coord.tileY),
			location: coord,
			sprite:   pixel.NewSprite(pic, pic.Bounds()),
		},
	}

	return obstacle
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

// func NewWinCondition() (winCondition WinCondition) {
// 	winCondition = WinCondition{object{25, 25, Coord{0, 0}}, 100}
// 	return winCondition
// }

// func NewLossCondition() (lossCondition LossCondition) {
// 	lossCondition = LossCondition{object{25, 25, Coord{0, 0}}, -100}
// 	return lossCondition
// }

package main

import (
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/pixelgl"
)

type object struct {
	// Generic object in the grid game
	// Can be a gopher, obstacle, reward,
	windowX  float64
	windowY  float64
	location Coord

	sprite *pixel.Sprite
	// objectType string

	// TODO: look into making separate structs for different objects
}

type Player struct {
	object

	score     float64
	direction float64 // TODO: improve; -1 = left, 1 = right

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
	mat := pixel.IM
	mat = mat.ScaledXY(pixel.ZV, pixel.V(p.direction, 1))
	mat = mat.Moved(pixel.V(p.windowX, p.windowY))
	p.sprite.Draw(win, mat)
}

func (o Obstacle) Draw(win *pixelgl.Window) {
	// TODO
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

// Generates an object of type Player
func NewPlayer(coord Coord) (player Player) {
	pic, err := loadPicture("assets/gopher50x50.png")
	if err != nil {
		panic(err)
	}

	player = Player{
		object: object{
			windowX: 25,
			windowY: 25,
			location: coord,
			sprite: pixel.NewSprite(pic, pic.Bounds()),
		},
		score:     0,
		direction: 1.0,
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
			windowX: 25,
			windowY: 25,
			location: coord,
			sprite: pixel.NewSprite(pic, pic.Bounds()),
		},
	}

	return obstacle
}

// func NewWinCondition() (winCondition WinCondition) {
// 	winCondition = WinCondition{object{25, 25, Coord{0, 0}}, 100}
// 	return winCondition
// }

// func NewLossCondition() (lossCondition LossCondition) {
// 	lossCondition = LossCondition{object{25, 25, Coord{0, 0}}, -100}
// 	return lossCondition
// }

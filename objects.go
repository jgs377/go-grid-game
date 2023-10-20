package main

import "github.com/gopxl/pixel/v2"

type object struct {
	// Generic object in the grid game
	// Can be a gopher, obstacle, reward,
	windowX  float64
	windowY  float64
	location Coord

	// objectType string

	// TODO: look into making separate structs for different objects
}

type Player struct {
	object

	score     int
	direction float64 // TODO: improve; -1 = left, 1 = right

	sprite *pixel.Sprite
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
	Draw()
}

func (p Player) Draw() {
	// TODO
}

func (o Obstacle) Draw() {
	// TODO
}

func (r Reward) Draw() {
	// TODO
}

func (w WinCondition) Draw() {
	// TODO
}

func (l LossCondition) Draw() {
	// TODO
}

// Generates an object of type Player
func NewPlayer() (player Player) {
	pic, err := loadPicture("assets/gopher50x50.png")
	if err != nil {
		panic(err)
	}

	player = Player{
		object: object{
			windowX: 25, 
			windowY: 25, 
			location: Coord{
				tileX: 0,
				tileY: 0,
			},
		}, 
		score: 0, 
		direction: 1,
		sprite: pixel.NewSprite(pic, pic.Bounds()),
	}

	return player
}

func NewObstacle() (obstacle Obstacle) {
	obstacle = Obstacle{object{25, 25, Coord{0, 0}}}
	return obstacle
}

func NewWinCondition() (winCondition WinCondition) {
	winCondition = WinCondition{object{25, 25, Coord{0, 0}}, 100}
	return winCondition
}

func NewLossCondition() (lossCondition LossCondition) {
	lossCondition = LossCondition{object{25, 25, Coord{0, 0}}, -100}
	return lossCondition
}

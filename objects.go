package main

type object struct {
	// Generic object in the grid game
	// Can be a gopher, obstacle, reward,
	windowX  float64
	windowY  float64
	location Coord

	objectType string

	// TODO: look into making separate structs for different objects
	direction float64 // TODO: improve; -1 = left, 1 = right
}

type Player struct {
	*object

	score int
}

type Obstacle struct {
	*object

	// In case we allow breaking obstacles in the future
	breakable    bool
	broken       bool
	breakPenalty float64
	walkPenalty  float64
}

type Reward struct {
	*object

	value int
}

type EndCondition struct {
	*object

	value int
}

type ObjectInterface interface {
	Draw()
}

func (obj *object) Draw() {
	// TODO
}

func (obj *object) loadAssets() {}

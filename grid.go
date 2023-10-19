package main

type Coord struct {
	tileX int
	tileY int
}

type Grid struct {
	sizeX int
	sizeY int
	tiles [][]ObjectInterface
}

func (g Grid) IsReward(coord Coord) bool {
	// TODO
	return true
}

func (g Grid) IsObstacle(coord Coord) bool {
	// TODO
	return true
}

func (g Grid) IsWinCondition(coord Coord) bool {
	// TODO
	return true
}

func (g Grid) IsLossCondition(coord Coord) bool {
	// TODO
	return true
}

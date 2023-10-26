package main

import "github.com/gopxl/pixel/v2/pixelgl"

const (
	North int = iota
	South
	East
	West
)

// Represents a coordinate pointing to a tile on the Grid
type Coord struct {
	tileX int
	tileY int
}

// Represents the X*Y grid on which the game takes place
type Grid struct {
	sizeX int
	sizeY int
	tiles [][]ObjectInterface
}

// Generates a new grid
func GenerateGrid(sizeX int, sizeY int, player *Player) (grid Grid) {
	grid.sizeX = sizeX
	grid.sizeY = sizeY

	// Make the 2D grid representation
	grid.tiles = make([][]ObjectInterface, sizeX)
	for i := range grid.tiles {
		grid.tiles[i] = make([]ObjectInterface, sizeY)
	}

	grid.tiles[0][0] = player
	grid.tiles[2][3] = NewObstacle(Coord{2, 3})
	grid.tiles[7][5] = NewObstacle(Coord{7, 5})
	grid.tiles[7][4] = NewObstacle(Coord{7, 4})
	grid.tiles[8][8] = NewReward(Coord{8, 8})

	return grid
}

// Determines if there is an object of type Reward on the tile at location (coord.X, coord.Y)
func (g Grid) IsReward(coord Coord) (isReward bool) {
	toTest := g.tiles[coord.tileX][coord.tileY]
	_, isReward = toTest.(Reward)
	return isReward
}

// Determines if there is an object of type Obstacle on the tile at location (coord.X, coord.Y)
func (g Grid) IsObstacle(coord Coord) (isObstacle bool) {
	toTest := g.tiles[coord.tileX][coord.tileY]
	_, isObstacle = toTest.(Obstacle)
	return isObstacle
}

// Determines if there is an object of type WinCondition on the tile at location (coord.X, coord.Y)
func (g Grid) IsWinCondition(coord Coord) (isWinCondition bool) {
	toTest := g.tiles[coord.tileX][coord.tileY]
	_, isWinCondition = toTest.(WinCondition)
	return isWinCondition
}

// Determines if there is an object of type EndCondition on the tile at location (coord.X, coord.Y)
func (g Grid) IsLossCondition(coord Coord) (isLossCondition bool) {
	toTest := g.tiles[coord.tileX][coord.tileY]
	_, isLossCondition = toTest.(LossCondition)
	return isLossCondition
}

// Determines if the tile a location (coord.X, coord.Y) is in bounds
func (g Grid) IsValidTile(coord Coord) (isValidTile bool) {
	if coord.tileX < 0 || coord.tileY < 0 {
		return false
	}
	if coord.tileX >= g.sizeX || coord.tileY >= g.sizeY {
		return false
	}
	if g.IsObstacle(coord) {
		return false
	}
	return true
}

func (c Coord) Shift(direction int) (shiftedCoord Coord) {
	switch direction {
	case North:
		return Coord{c.tileX, c.tileY + 1}
	case South:
		return Coord{c.tileX, c.tileY - 1}
	case East:
		return Coord{c.tileX - 1, c.tileY}
	case West:
		return Coord{c.tileX + 1, c.tileY}
	default:
		return c
	}
}

func (g *Grid) Draw(win *pixelgl.Window) {
	for x := 0; x < g.sizeX; x++ {
		for y := 0; y < g.sizeY; y++ {
			if g.tiles[x][y] != nil {
				g.tiles[x][y].Draw(win)
			}
		}
	}
}

package main

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
	return true
}

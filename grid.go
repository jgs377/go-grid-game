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

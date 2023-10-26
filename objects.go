package main

import (
	"fmt"
	"image"
	"os"

	_ "image/png"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/pixelgl"
)

const (
	debug = false
	asset_path = "assets/"
	playerPath = asset_path + "gopher50x50.png"
	playerBackPath = asset_path + "gopher50x50back.png"
	playerFrontPath = asset_path + "gopher50x50front.png"
	rewardPath = asset_path + "carrot.png"
	winConditionPath = asset_path + "dollar-bag2.png"
	obstaclePath = asset_path + "rock.png"
	lossConditionPath = asset_path + "hole.png"
)

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

type EndCondition struct {
	object
	value int
}

// All object types present in the game implement this ObjectInterface
type ObjectInterface interface {
	Draw(*pixelgl.Window)
}

func (o object) Draw(win *pixelgl.Window) {
	mat := pixel.IM.Moved(pixel.V(o.windowX, o.windowY))
	o.sprite.Draw(win, mat)
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
	o.object.Draw(win)
}

func (r Reward) Draw(win *pixelgl.Window) {
	r.object.Draw(win)
}

func (e EndCondition) Draw(win *pixelgl.Window) {
	e.object.Draw(win)
}

var movement = map[int]Coord{
	North: {0, 1},
	South: {0, -1},
	East: {-1, 0},
	West: {1, 0},
}

// Moves the player in the grid
func (p *Player) Move(direction int, grid *Grid) {
	move := movement[direction]
	newX := p.location.tileX + move.tileX
	newY := p.location.tileY + move.tileY

	if grid.IsReward(Coord{newX, newY}) {
		p.score += float64(grid.tiles[newX][newY].(Reward).value)
	}

	if grid.IsEndCondition(Coord{newX, newY}) {
		p.score += float64(grid.tiles[newX][newY].(EndCondition).value)
		grid.gameOver = true
	}

	grid.tiles[newX][newY] = grid.tiles[p.location.tileX][p.location.tileY]
    grid.tiles[p.location.tileX][p.location.tileY] = nil

	p.location.tileX = newX
	p.location.tileY = newY
	p.windowX += float64(50 * move.tileX)
	p.windowY += float64(50 * move.tileY)
	p.score -= 0.1

	if debug {
		for _, j := range grid.tiles {
			fmt.Println(j)
		}
		fmt.Println("================")
	}
}

// Generates an object of type Player
func NewPlayer(coord Coord) (player Player) {
	sprite1 := newSpriteFromPath(playerPath)
	sprite2 := newSpriteFromPath(playerBackPath)
	sprite3 := newSpriteFromPath(playerFrontPath)
	player = Player{
		object: object{
			windowX:  25,
			windowY:  25,
			location: coord,
			sprite:   sprite1,
		},
		score:       0,
		direction:   1.0,
		backSprite:  sprite2,
		frontSprite: sprite3,
	}

	return player
}

// Generates an object of type Obstacle
func NewObstacle(coord Coord) (obstacle Obstacle) {
	sprite := newSpriteFromPath(obstaclePath)

	obstacle = Obstacle{
		object: object{
			windowX:  float64(25 + 50*coord.tileX),
			windowY:  float64(25 + 50*coord.tileY),
			location: coord,
			sprite:   sprite,
		},
	}

	return obstacle
}

// Generates an object of type Reward
func NewReward(coord Coord) (reward Reward) {
	sprite := newSpriteFromPath(rewardPath)
	
	reward = Reward{
		object: object{
			windowX: float64(25+50*coord.tileX),
			windowY: float64(25+50*coord.tileY),
			location: coord,
			sprite: sprite,
		},
		value: 5,
	}

	return reward
}

func NewEndCondition(coord Coord, value int) (endCondition EndCondition) {
	var sprite *pixel.Sprite

	if value > 0 {
		sprite = newSpriteFromPath(winConditionPath)
	} else {
		sprite = newSpriteFromPath(lossConditionPath)
	}

	endCondition = EndCondition{
		object: object{
			windowX: float64(25+50*coord.tileX),
			windowY: float64(25+50*coord.tileY),
			location: coord,
			sprite: sprite,
		},
		value: value,
	}
	return endCondition
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

func newSpriteFromPath(path string) (sprite *pixel.Sprite) {
	pic, err := loadPicture(path)
	if err != nil {
		panic(err)
	}
	return pixel.NewSprite(pic, pic.Bounds())
}


// func NewLossCondition() (lossCondition LossCondition) {
// 	lossCondition = LossCondition{object{25, 25, Coord{0, 0}}, -100}
// 	return lossCondition
// }

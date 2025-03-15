package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/heschmat/go-ebiten-asteroids/assets"
)

const (
	rps 			= math.Pi // rotation per second
	maxAcceleration = 8
	ScreenWidth		= 1280
	ScreenHeight	= 720 // 16x9 ratio (w, h)
)

type Player struct{
	game		*Game
	sprite 		*ebiten.Image
	rotation	float64
	pos			Vector  // <x, y>
	Velo		float64 // velocity
}

func NewPlayer(game *Game) *Player {
	sprite := assets.PlayerSprite

	// Create player on center of the screen.
	bounds := sprite.Bounds()
	playerInitialPos := Vector{
		X: ScreenWidth/2 - float64(bounds.Dx()) / 2,
		Y: ScreenHeight/2 - float64(bounds.Dy()) / 2,
	}

	p := &Player{
		sprite: sprite,
		game: 	game,
		pos: 	playerInitialPos,
	}

	return p
}

func (p *Player) Draw(screen *ebiten.Image) {
	bounds := p.sprite.Bounds()

	halfWidth := float64(bounds.Dx()) / 2
	halfHeight := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(-halfWidth, -halfHeight)
	op.GeoM.Rotate(p.rotation)
	op.GeoM.Translate(halfWidth, halfHeight)

	op.GeoM.Translate(p.pos.X, p.pos.Y)

	screen.DrawImage(p.sprite, op)
}

func (p *Player) Update() {
	speed := rps / float64(ebiten.TPS()) // text per second

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		// fmt.Println('L')
		p.rotation -= speed
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.rotation += speed
	}

	p.accelerate()
}

func (p *Player) accelerate() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.keepOnScreen()

		p.Velo = math.Max(p.Velo + 4, maxAcceleration)

		p.pos.X += math.Sin(p.rotation) * p.Velo
		p.pos.Y -= math.Cos(p.rotation) * p.Velo
	}
}

func (p *Player) keepOnScreen() {
	if p.pos.X >= float64(ScreenWidth) {
		p.pos.X = 0
	} else if p.pos.X < 0 {
		p.pos.X = ScreenWidth
	}

	if p.pos.Y >= float64(ScreenHeight) {
		p.pos.Y = 0
	} else if p.pos.Y < 0 {
		p.pos.Y = ScreenHeight
	}
}

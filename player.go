package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/heschmat/go-ebiten-asteroids/assets"
)

const rps = math.Pi // rotation per second

type Player struct{
	sprite 		*ebiten.Image
	rotation	float64
}

func NewPlayer(game *Game) *Player {
	sprite := assets.PlayerSprite

	p := &Player{
		sprite: sprite,
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
}

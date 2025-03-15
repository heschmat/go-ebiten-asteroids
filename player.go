package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/heschmat/go-ebiten-asteroids/assets"
)

type Player struct{
	sprite *ebiten.Image
}

func NewPlayer(game *Game) *Player {
	sprite := assets.PlayerSprite

	p := &Player{
		sprite: sprite,
	}

	return p
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	screen.DrawImage(p.sprite, op)
}

func (p *Player) Update() {

}

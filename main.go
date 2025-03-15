package main

import "github.com/hajimehoshi/ebiten/v2"

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(w, h int) (ScreenWidth, ScreenHeight int) {
	return w, h
}

func main() {
	g := &Game{}

	err := ebiten.RunGame(g)

	if err != nil {
		panic(err) // for now
	}
}

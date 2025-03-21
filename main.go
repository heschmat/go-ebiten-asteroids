package main

import "github.com/hajimehoshi/ebiten/v2"

type Game struct{
	player *Player
}

func (g *Game) Update() error {
	g.player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

}

func (g *Game) Layout(w, h int) (ScreenWidth, ScreenHeight int) {
	return w, h
}

func main() {
	g := &Game{}
	g.player = NewPlayer(g)

	ebiten.SetWindowTitle("Go Asteroids!")
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)

	err := ebiten.RunGame(g)

	if err != nil {
		panic(err) // for now
	}
}

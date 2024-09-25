package main

import (
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/ebitick"
)

type Game struct {
	TimerSystem *ebitick.TimerSystem
	arena       *map[Position]Square
	current     *Tetromino
}

func newGame() *Game {
	game := &Game{
		TimerSystem: ebitick.NewTimerSystem(),
		arena:       generateBorders(),
	}
	timer := game.TimerSystem.After(time.Second, func() {
		game.current.Descend()
	})
	timer.Loop = true
	return game
}

func (g *Game) Update() error {
	g.TimerSystem.Update()
	if g.current == nil {
		g.current = newTetromino(color.RGBA{255, 0, 0, 255})
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, square := range *g.arena {
		img, opts := square.Image(color.RGBA{120, 120, 120, 255})
		screen.DrawImage(img, &opts)
	}
	for _, part := range g.current.Parts {
		img, opts := newSquare(part).Image(g.current.Color)
		screen.DrawImage(img, &opts)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 880
}

func main() {
	ebiten.SetWindowSize(800, 880)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(newGame()); err != nil {
		log.Fatal(err)
	}
}

func generateBorders() *map[Position]Square {
	borders := make(map[Position]Square)
	for y := range 22 {
		posLeft := Position{X: 1, Y: y + 1}
		borders[posLeft] = *newSquare(posLeft)
		posRight := Position{X: 12, Y: y + 1}
		borders[posRight] = *newSquare(posRight)
	}
	for x := range 10 {
		posTop := Position{X: x + 2, Y: 1}
		borders[posTop] = *newSquare(posTop)
		posBot := Position{X: x + 2, Y: 22}
		borders[posBot] = *newSquare(posBot)
	}
	return &borders
}

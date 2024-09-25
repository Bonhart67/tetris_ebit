package main

import (
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	input "github.com/quasilyte/ebitengine-input"
	"github.com/solarlune/ebitick"
)

const (
	UpdateTime                  = 200
	ActionMoveLeft input.Action = iota
	ActionMoveRight
)

type Game struct {
	TimerSystem *ebitick.TimerSystem
	arena       *map[Position]Square
	current     *Tetromino
	inputSystem input.System
  input *input.Handler
}

func newGame() *Game {
	game := &Game{
		TimerSystem: ebitick.NewTimerSystem(),
		arena:       generateBorders(),
	}
	game.inputSystem.Init(input.SystemConfig{DevicesEnabled: input.AnyDevice})

	keymap := input.Keymap{
		ActionMoveLeft:  {input.KeyGamepadLeft, input.KeyLeft, input.KeyA},
		ActionMoveRight: {input.KeyGamepadRight, input.KeyRight, input.KeyD},
	}

  game.input = game.inputSystem.NewHandler(0, keymap)

	timer := game.TimerSystem.After(time.Millisecond*UpdateTime, func() {
		if game.current != nil && !game.current.IsStuck(game.arena) {
			game.current.Descend()
		} else {
			for _, part := range *game.current.Parts {
				(*game.arena)[part] = *newSquare(part)
			}
			game.current = nil
		}
	})
	timer.Loop = true
	return game
}

func (g *Game) Update() error {
	g.TimerSystem.Update()
	if g.current == nil {
		g.current = newTetromino(color.RGBA{255, 0, 0, 255})
	}
  if g.input.ActionIsJustPressed(ActionMoveLeft) {
    g.current.Move(-1)
  }
  if g.input.ActionIsJustPressed(ActionMoveRight) {
    g.current.Move(1)
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

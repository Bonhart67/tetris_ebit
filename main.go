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
	arena       *Arena
	current     *Tetromino
	inputSystem input.System
	input       *input.Handler
}

func newGame() *Game {
	game := &Game{
		TimerSystem: ebitick.NewTimerSystem(),
		arena:       newArena(),
	}
	game.inputSystem.Init(input.SystemConfig{DevicesEnabled: input.AnyDevice})

	keymap := input.Keymap{
		ActionMoveLeft:  {input.KeyGamepadLeft, input.KeyLeft, input.KeyA},
		ActionMoveRight: {input.KeyGamepadRight, input.KeyRight, input.KeyD},
	}

	game.input = game.inputSystem.NewHandler(0, keymap)

	timer := game.TimerSystem.After(time.Millisecond*UpdateTime, func() {
		if game.current != nil && game.current.CanMoveDown(game.arena) {
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
		g.current.MoveLeft(g.arena)
	}
	if g.input.ActionIsJustPressed(ActionMoveRight) {
		g.current.MoveRight(g.arena)
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

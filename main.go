package main

import (
	"errors"
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
	ActionRotate
	ActionCloseGame
)

type Game struct {
	TimerSystem *ebitick.TimerSystem
	arena       *Arena
	current     *Tetromino
	inputSystem input.System
	input       *input.Handler
	over        bool
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
		ActionRotate:    {input.KeyGamepadDown, input.KeyDown, input.KeyS},
		ActionCloseGame: {input.KeyEscape, input.KeyQ},
	}

	game.input = game.inputSystem.NewHandler(0, keymap)

	timer := game.TimerSystem.After(time.Millisecond*UpdateTime, func() {
		if game.current != nil && game.current.Collides(game.arena) {
			game.over = true
		} else if game.current != nil && game.current.CanMoveDown(game.arena) {
			game.current.Descend()
		} else {
			game.arena.Add(*game.current)
			game.current = nil
		}
	})
	timer.Loop = true
	return game
}

func (g *Game) Update() error {
	g.TimerSystem.Update()
	if !g.over && g.current == nil {
		g.current = newTetromino()
	}
	if g.input.ActionIsJustPressed(ActionMoveLeft) {
		g.current.MoveLeft(g.arena)
	}
	if g.input.ActionIsJustPressed(ActionMoveRight) {
		g.current.MoveRight(g.arena)
	}
	if g.input.ActionIsJustPressed(ActionRotate) {
		g.current.Rotate(g.arena)
	}
	if g.input.ActionIsJustPressed(ActionCloseGame) {
		return Terminated
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, square := range *g.arena {
		img, opts := square.Image()
		screen.DrawImage(img, &opts)
	}
	for _, part := range g.current.parts() {
		img, opts := newSquare(part, g.current.Color).Image()
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
		if err == Terminated {
			return
		}
		log.Fatal(err)
	}
}

var Terminated = errors.New("terminated")

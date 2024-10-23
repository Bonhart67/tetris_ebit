package main

import (
	"image/color"
	"math/rand"
)

func randomColor() color.Color {
	colors := []color.Color{
		color.RGBA{255, 0, 0, 255},
		color.RGBA{0, 255, 0, 255},
		color.RGBA{0, 255, 0, 255},
		color.RGBA{0, 128, 128, 255},
		color.RGBA{255, 165, 0, 255},
		color.RGBA{255, 255, 0, 255},
		color.RGBA{160, 32, 240, 255},
	}
	return colors[rand.Intn(len(colors))]
}

type Tetromino struct {
	Position
	Shape
	State int
	Color color.Color
}

func newTetromino() *Tetromino {
	return &Tetromino{
		Position: Position{X: 4, Y: 1},
		Shape:    &I{},
		State:    0,
		Color:    randomColor(),
	}
}

func (t *Tetromino) Collides(a *Arena) bool {
	for _, p := range t.parts() {
		if a.Contains(p.X, p.Y) {
			return true
		}
	}
	return false
}

func (t *Tetromino) parts() []Position {
	parts := t.Shape.parts(t.State)
	for i := range parts {
		parts[i] = Position{X: parts[i].X + t.Position.X, Y: parts[i].Y + t.Position.Y}
	}
	return parts
}

func (t *Tetromino) Descend() {
	t.Position.Y += 1
}

func (t *Tetromino) MoveLeft(a *Arena) {
	if t.canMoveLeft(a) {
		t.Position.X -= 1
	}
}

func (t *Tetromino) MoveRight(a *Arena) {
	if t.canMoveRight(a) {
		t.Position.X += 1
	}
}

func (t *Tetromino) Rotate(a *Arena) {
	prevState := t.State
	t.State = (t.State + 1) % 4
	if t.Collides(a) {
		t.State = prevState
	}
}

func (t *Tetromino) CanMoveDown(a *Arena) bool {
	bottom := make(map[int]Position)
	for _, part := range t.parts() {
		if val, ok := bottom[part.X]; !ok || (ok && val.Y < part.Y) {
			bottom[part.X] = part
		}
	}
	for _, part := range bottom {
		if a.Contains(part.X, part.Y+1) {
			return false
		}
	}
	return true
}

func (t *Tetromino) canMoveLeft(a *Arena) bool {
	left := make(map[int]Position)
	for _, part := range t.parts() {
		if present, ok := left[part.Y]; !ok || (ok && present.X > part.X) {
			left[part.Y] = part
		}
	}
	for _, part := range left {
		if a.Contains(part.X-1, part.Y) {
			return false
		}
	}
	return true
}

func (t *Tetromino) canMoveRight(a *Arena) bool {
	right := make(map[int]Position)
	for _, part := range t.parts() {
		if present, ok := right[part.Y]; !ok || (ok && present.X < part.X) {
			right[part.Y] = part
		}
	}
	for _, part := range right {
		if a.Contains(part.X+1, part.Y) {
			return false
		}
	}
	return true
}

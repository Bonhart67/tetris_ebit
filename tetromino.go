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

func randomParts() *[4]Position {
	tetrominos := [][4]Position{
		{{X: 6, Y: 2}, {X: 6, Y: 3}, {X: 6, Y: 4}, {X: 6, Y: 5}},
		{{X: 7, Y: 2}, {X: 7, Y: 3}, {X: 7, Y: 4}, {X: 6, Y: 4}},
		{{X: 6, Y: 2}, {X: 6, Y: 3}, {X: 6, Y: 4}, {X: 7, Y: 4}},
		{{X: 6, Y: 2}, {X: 6, Y: 3}, {X: 7, Y: 2}, {X: 7, Y: 3}},
		{{X: 5, Y: 3}, {X: 6, Y: 3}, {X: 6, Y: 2}, {X: 7, Y: 2}},
		{{X: 5, Y: 2}, {X: 6, Y: 2}, {X: 6, Y: 3}, {X: 7, Y: 2}},
		{{X: 5, Y: 2}, {X: 6, Y: 3}, {X: 6, Y: 2}, {X: 7, Y: 3}},
	}
	return &tetrominos[rand.Intn(len(tetrominos))]
}

type Tetromino struct {
	Parts *[4]Position
	Color color.Color
}

func newTetromino() *Tetromino {
	return &Tetromino{
		Parts: randomParts(),
		Color: randomColor(),
	}
}

func (t *Tetromino) Descend() {
	for i := range t.Parts {
		t.Parts[i].Y += 1
	}
}

func (t *Tetromino) MoveLeft(a *Arena) {
	if t.canMoveLeft(a) {
		for i := range t.Parts {
			t.Parts[i].X -= 1
		}
	}
}

func (t *Tetromino) MoveRight(a *Arena) {
	if t.canMoveRight(a) {
		for i := range t.Parts {
			t.Parts[i].X += 1
		}
	}
}

func (t *Tetromino) Rotate(a *Arena) {
	// TODO
	println("Rotate")
}

func (t *Tetromino) CanMoveDown(a *Arena) bool {
	bottom := make(map[int]Position)
	for _, part := range t.Parts {
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
	for _, part := range t.Parts {
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
	for _, part := range t.Parts {
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

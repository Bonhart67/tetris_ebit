package main

import (
	"image/color"
)

type Tetromino struct {
	Parts *[4]Position
	Color color.Color
}

func newTetromino(c color.Color) *Tetromino {
	return &Tetromino{
		Parts: &[4]Position{
			{X: 7, Y: 2},
			{X: 7, Y: 3},
			{X: 6, Y: 3},
			{X: 6, Y: 4},
		},
		Color: c,
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

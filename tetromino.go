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

func (t *Tetromino) IsStuck(arena *map[Position]Square) bool {
	bottom := make(map[int]Position)
	for _, part := range t.Parts {
		if val, ok := bottom[part.X]; !ok {
			bottom[part.X] = part
		} else if ok && val.Y < part.Y {
			bottom[part.X] = part
		}
	}
	for _, part := range bottom {
		if containsPosition(part.X, part.Y+1, arena) {
			return true
		}
	}
	return false
}

func containsPosition(x, y int, arena *map[Position]Square) bool {
	for _, p := range *arena {
		if p.X == x && p.Y == y {
			return true
		}
	}
	return false
}

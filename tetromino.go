package main

import "image/color"

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

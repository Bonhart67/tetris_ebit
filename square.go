package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Square struct {
	Position
	color color.Color
}

func newSquare(p Position, c color.Color) *Square {
	return &Square{
		Position: p,
		color:    c,
	}
}

type Position struct {
	X, Y int
}

func (s Square) Image() (*ebiten.Image, ebiten.DrawImageOptions) {
	img := ebiten.NewImage(40, 40)
	img.Fill(s.color)
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(s.screenPosition())
	return img, opts
}

func (s *Square) screenPosition() (x, y float64) {
	return float64(s.X-1) * 40, float64(s.Y-1) * 40
}

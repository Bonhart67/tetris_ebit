package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Square struct {
	Position
}

func newSquare(p Position) *Square {
	return &Square{
		Position: p,
	}
}

type Position struct {
	X, Y int
}

func (s Square) Image() (*ebiten.Image, ebiten.DrawImageOptions) {
	img := ebiten.NewImage(40, 40)
	img.Fill(color.RGBA{120, 120, 120, 255})
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(s.screenPosition())
	return img, opts
}

func (s *Square) screenPosition() (x, y float64) {
	return float64(s.X - 1) * 40, float64(s.Y - 1) * 40
}

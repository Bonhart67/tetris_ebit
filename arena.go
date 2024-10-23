package main

import "image/color"

type Arena map[Position]Square

func newArena() *Arena {
	return (*Arena)(generateBorders())
}

func (a *Arena) Contains(x, y int) bool {
	_, ok := a.Get(x, y)
	return ok
}

func (a *Arena) Get(x, y int) (val *Square, ok bool) {
	for _, sq := range *a {
		if sq.X == x && sq.Y == y {
			return &sq, true
		}
	}
	return nil, false
}

func (a Arena) Add(t Tetromino) {
	for _, part := range t.parts() {
		a[part] = *newSquare(part, t.Color)
	}
}

func generateBorders() *map[Position]Square {
	borders := make(map[Position]Square)
	for y := range 22 {
		posLeft := Position{X: 0, Y: y}
		borders[posLeft] = *newSquare(posLeft, color.RGBA{120, 120, 120, 255})
		posRight := Position{X: 11, Y: y}
		borders[posRight] = *newSquare(posRight, color.RGBA{120, 120, 120, 255})
	}
	for x := range 10 {
		posTop := Position{X: x + 1, Y: 0}
		borders[posTop] = *newSquare(posTop, color.RGBA{120, 120, 120, 255})
		posBot := Position{X: x + 1, Y: 21}
		borders[posBot] = *newSquare(posBot, color.RGBA{120, 120, 120, 255})
	}
	return &borders
}

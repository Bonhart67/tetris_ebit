package main

import (
	"image/color"
)

type Arena struct {
	squares map[Position]Square
}

func newArena() Arena {
	return generateBorders()
}

func (a Arena) Contains(x, y int) bool {
	_, ok := a.Get(x, y)
	return ok
}

func (a Arena) Get(x, y int) (v Square, ok bool) {
	for _, sq := range a.squares {
		if sq.X == x && sq.Y == y {
			return sq, true
		}
	}
	return Square{}, false
}

func (a *Arena) Delete(x, y int) {
	for k, v := range a.squares {
		if v.X == x && v.Y == y {
			delete(a.squares, k)
		}
	}
}

func (a *Arena) Add(t Tetromino) {
	for _, part := range t.parts() {
		a.squares[part] = *newSquare(part, t.Color)
	}
	for y := 20; y > 1; y-- {
		if a.countInRow(y) == 10 {
			a.deleteRow(y)
			a.pushDownOnce(y)
			// y++
		}
	}
}

func (a Arena) countInRow(y int) int {
	count := -2 // left and right border
	for k := range a.squares {
		if k.Y == y {
			count++
		}
	}
	return count
}

func (a *Arena) deleteRow(y int) {
	for k := range a.squares {
		if k.Y == y && k.X != 0 && k.X != 11 {
			delete(a.squares, k)
		}
	}
}

func (a *Arena) pushDownOnce(y int) {
	for current := y - 1; current > 1; current-- {
		for x := 1; x < 11; x++ {
			if s, ok := a.Get(x, current); ok {
				nextPos := Position{X: x, Y: current + 1}
				a.squares[nextPos] = *newSquare(nextPos, s.color)
				a.Delete(x, current)
			}
		}
	}
}

func generateBorders() Arena {
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
	return Arena{squares: borders}
}

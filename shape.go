package main

type Shape interface {
	parts(i int) []Position
}

type I struct{}

func (s I) parts(i int) []Position {
	variants := [][]Position{
		{{X: 0, Y: 1}, {X: 1, Y: 1}, {X: 2, Y: 1}, {X: 3, Y: 1}},
		{{X: 2, Y: 0}, {X: 2, Y: 1}, {X: 2, Y: 2}, {X: 2, Y: 3}},
		{{X: 0, Y: 2}, {X: 1, Y: 2}, {X: 2, Y: 2}, {X: 3, Y: 2}},
		{{X: 1, Y: 0}, {X: 1, Y: 1}, {X: 1, Y: 2}, {X: 1, Y: 3}},
	}
	return variants[i]
}

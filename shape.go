package main

import "math/rand"

type Shape interface {
	parts(i int) []Position
}

func RandomShape() Shape {
	shapes := []Shape{&I{}, &O{}, &T{}, &S{}, &Z{}, &J{}, &L{}}
	return shapes[rand.Intn(len(shapes))]
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

type O struct{}

func (s O) parts(i int) []Position {
	variants := [][]Position{
		{{X: 1, Y: 1}, {X: 1, Y: 2}, {X: 2, Y: 1}, {X: 2, Y: 2}},
		{{X: 1, Y: 1}, {X: 1, Y: 2}, {X: 2, Y: 1}, {X: 2, Y: 2}},
		{{X: 1, Y: 1}, {X: 1, Y: 2}, {X: 2, Y: 1}, {X: 2, Y: 2}},
		{{X: 1, Y: 1}, {X: 1, Y: 2}, {X: 2, Y: 1}, {X: 2, Y: 2}},
	}
	return variants[i]
}

type T struct{}

func (s T) parts(i int) []Position {
	variants := [][]Position{
		{{X: 1, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}, {X: 2, Y: 1}},
		{{X: 1, Y: 0}, {X: 1, Y: 1}, {X: 1, Y: 2}, {X: 2, Y: 1}},
		{{X: 0, Y: 1}, {X: 1, Y: 1}, {X: 2, Y: 1}, {X: 1, Y: 2}},
		{{X: 1, Y: 0}, {X: 1, Y: 1}, {X: 1, Y: 2}, {X: 0, Y: 1}},
	}
	return variants[i]
}

type S struct{}

func (s S) parts(i int) []Position {
	variants := [][]Position{
		{{X: 1, Y: 0}, {X: 2, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}},
		{{X: 1, Y: 0}, {X: 1, Y: 1}, {X: 2, Y: 1}, {X: 2, Y: 2}},
		{{X: 1, Y: 1}, {X: 2, Y: 1}, {X: 0, Y: 2}, {X: 1, Y: 2}},
		{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}, {X: 1, Y: 2}},
	}
	return variants[i]
}

type Z struct{}

func (s Z) parts(i int) []Position {
	variants := [][]Position{
		{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 1, Y: 1}, {X: 2, Y: 1}},
		{{X: 2, Y: 0}, {X: 2, Y: 1}, {X: 1, Y: 1}, {X: 1, Y: 2}},
		{{X: 0, Y: 1}, {X: 1, Y: 1}, {X: 1, Y: 2}, {X: 2, Y: 2}},
		{{X: 1, Y: 0}, {X: 1, Y: 1}, {X: 0, Y: 1}, {X: 0, Y: 2}},
	}
	return variants[i]
}

type J struct{}

func (s J) parts(i int) []Position {
	variants := [][]Position{
		{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}, {X: 2, Y: 1}},
		{{X: 1, Y: 0}, {X: 1, Y: 1}, {X: 1, Y: 2}, {X: 2, Y: 0}},
		{{X: 0, Y: 1}, {X: 1, Y: 1}, {X: 2, Y: 1}, {X: 2, Y: 2}},
		{{X: 1, Y: 0}, {X: 1, Y: 1}, {X: 1, Y: 2}, {X: 0, Y: 1}},
	}
	return variants[i]
}

type L struct{}

func (s L) parts(i int) []Position {
	variants := [][]Position{
		{{X: 0, Y: 1}, {X: 1, Y: 1}, {X: 2, Y: 1}, {X: 2, Y: 0}},
		{{X: 1, Y: 0}, {X: 1, Y: 1}, {X: 1, Y: 2}, {X: 2, Y: 2}},
		{{X: 0, Y: 1}, {X: 1, Y: 1}, {X: 2, Y: 1}, {X: 0, Y: 2}},
		{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 1, Y: 1}, {X: 1, Y: 2}},
	}
	return variants[i]
}

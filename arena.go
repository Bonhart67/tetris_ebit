package main

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

func generateBorders() *map[Position]Square {
	borders := make(map[Position]Square)
	for y := range 22 {
		posLeft := Position{X: 1, Y: y + 1}
		borders[posLeft] = *newSquare(posLeft)
		posRight := Position{X: 12, Y: y + 1}
		borders[posRight] = *newSquare(posRight)
	}
	for x := range 10 {
		posTop := Position{X: x + 2, Y: 1}
		borders[posTop] = *newSquare(posTop)
		posBot := Position{X: x + 2, Y: 22}
		borders[posBot] = *newSquare(posBot)
	}
	return &borders
}

package core

import "math/rand"

func sample(cells []*Cell) *Cell {
	return cells[rand.Intn(len(cells))]
}

func coinflip() bool {
	return rand.Intn(2) == 0
}

package core

import "math/rand"

func (g *Grid) BinaryTreeMaze() {
	g.eachCell(func(c *Cell) {
		neighbors := make([]*Cell, 0, 2)
		if c.north != nil {
			neighbors = append(neighbors, c.north)
		}
		if c.east != nil {
			neighbors = append(neighbors, c.east)
		}
		if len(neighbors) > 0 {
			index := rand.Intn(len(neighbors))
			neighbor := neighbors[index]
			c.link(neighbor)
		}
	})
}

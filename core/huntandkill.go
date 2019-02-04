package core

func filter(vs []*Cell, f func(*Cell) bool) []*Cell {
	vsf := make([]*Cell, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func (g *Grid) HuntAndKillMaze() {
	current := g.randomCell()

	for current != nil {
		unvisitedNeighbors := filter(current.neighbors(), func(c *Cell) bool {
			return len(c.links) == 0
		})

		if len(unvisitedNeighbors) > 0 {
			neighbor := sample(unvisitedNeighbors)
			current.link(neighbor)
			current = neighbor
		} else {
			current = nil

			g.eachCell(func(c *Cell) {
				visitedNeighbors := filter(c.neighbors(), func(n *Cell) bool {
					return len(n.links) > 0
				})
				if len(c.links) == 0 && len(visitedNeighbors) > 0 {
					current = c
					neighbor := sample(visitedNeighbors)
					current.link(neighbor)
				}
			})
		}
	}
}

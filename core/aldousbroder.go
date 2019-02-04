package core

func (g *Grid) AldousBroderMaze() {
	cell := g.randomCell()
	unvisited := g.size() - 1

	for unvisited > 0 {
		neighbor := sample(cell.neighbors())

		if len(neighbor.links) == 0 {
			cell.link(neighbor)
			unvisited--
		}

		cell = neighbor
	}
}

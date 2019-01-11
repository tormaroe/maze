package core

func (g *Grid) SidewinderMaze() {
	g.eachRow(func(row []*Cell) {
		run := make([]*Cell, 0, len(row))

		for _, c := range row {
			run = append(run, c)
			atEastBound := c.east == nil
			atNorthBound := c.north == nil
			shouldClose := atEastBound || (!atNorthBound && coinflip())

			if shouldClose {
				member := sample(run)
				if member.north != nil {
					member.link(member.north)
				}
				run = make([]*Cell, 0, len(row))
			} else {
				c.link(c.east)
			}
		}
	})
}

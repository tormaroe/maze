package core

type Cell struct {
	row    int
	column int
	north  *Cell
	south  *Cell
	east   *Cell
	west   *Cell
	links  map[*Cell]bool
}

func newCell(row, column int) *Cell {
	return &Cell{row: row, column: column, links: make(map[*Cell]bool)}
}

func (c *Cell) link(other *Cell) {
	c.links[other] = true
	other.links[c] = true
}

func (c *Cell) unlink(other *Cell) {
	delete(c.links, other)
	delete(other.links, c)
}

func (c *Cell) isLinked(other *Cell) bool {
	_, ok := c.links[other]
	return ok
}

func (c *Cell) linkedCells() (result []*Cell) {
	result = make([]*Cell, len(c.links))
	i := 0
	for k := range c.links {
		result[i] = k
		i++
	}
	return
}

func (c *Cell) neighbors() (list []*Cell) {
	list = make([]*Cell, 0, 4)
	if c.east != nil {
		list = append(list, c.east)
	}
	if c.west != nil {
		list = append(list, c.west)
	}
	if c.north != nil {
		list = append(list, c.north)
	}
	if c.south != nil {
		list = append(list, c.south)
	}
	return
}

package core

type Distance struct {
	root      *Cell
	cellDists map[*Cell]int
}

func NewDistance(root *Cell) Distance {
	return Distance{root: root,
		cellDists: map[*Cell]int{
			root: 0,
		}}
}

func (d Distance) distanceTo(cell *Cell) (int, bool) {
	v, ok := d.cellDists[cell]
	return v, ok
}

func (d *Distance) setDistanceTo(cell *Cell, v int) {
	d.cellDists[cell] = v
}

func (d *Distance) cells() (result []*Cell) {
	result = make([]*Cell, len(d.cellDists))
	i := 0
	for k := range d.cellDists {
		result[i] = k
		i++
	}
	return
}

func (c *Cell) Distances() *Distance {
	d := NewDistance(c)
	frontier := []*Cell{c}

	for len(frontier) > 0 {
		newFrontier := []*Cell{}

		for _, cell := range frontier {
			for _, linked := range cell.linkedCells() {
				if _, found := d.distanceTo(linked); found {
					continue
				}
				v, _ := d.distanceTo(cell)
				d.setDistanceTo(linked, v+1)
				newFrontier = append(newFrontier, linked)
			}
		}

		frontier = newFrontier
	}

	return &d
}

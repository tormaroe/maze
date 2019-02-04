package core

type Distance struct {
	root      *Cell
	cellDists map[*Cell]int
	maximum   int   // cache
	farthest  *Cell // cache
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

func (d Distance) PathToCell(goal *Cell) *Distance {
	current := goal

	breadcrumbs := NewDistance(d.root)
	v, _ := d.distanceTo(current)
	breadcrumbs.setDistanceTo(current, v)

	for current != d.root {
		for _, neighbor := range current.linkedCells() {
			vn, _ := d.distanceTo(neighbor)
			vc, _ := d.distanceTo(current)
			if vn < vc {
				breadcrumbs.setDistanceTo(neighbor, vn)
				current = neighbor
				break
			}
		}
	}

	return &breadcrumbs
}

func (d *Distance) Max() (*Cell, int) {
	if d.farthest == nil {
		d.maximum = 0
		d.farthest = d.root
		for cell, distance := range d.cellDists {
			if distance > d.maximum {
				d.farthest = cell
				d.maximum = distance
			}
		}
	}
	return d.farthest, d.maximum
}

func (g Grid) LongestPath() *Distance {
	start := g.CellAt(0, 0)
	distances := start.Distances()
	newStart, _ := distances.Max()
	newDistances := newStart.Distances()
	goal, _ := newDistances.Max()
	return newDistances.PathToCell(goal)
}

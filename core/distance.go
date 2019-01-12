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

func (d Distance) Max() (maxCell *Cell, maxDistance int) {
	maxDistance = 0
	maxCell = d.root
	for cell, distance := range d.cellDists {
		if distance > maxDistance {
			maxCell = cell
			maxDistance = distance
		}
	}
	return
}

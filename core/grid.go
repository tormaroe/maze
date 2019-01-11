package core

import (
	"math/rand"
)

type Grid struct {
	rows    int
	columns int
	grid    [][]*Cell
}

func NewGrid(rows, columns int) (g Grid) {
	g = Grid{rows: rows, columns: columns}
	g.grid = make([][]*Cell, rows)
	for i := range g.grid {
		g.grid[i] = make([]*Cell, columns)
		for j := 0; j < columns; j++ {
			g.grid[i][j] = newCell(i, j)
		}
	}
	g.eachCell(func(c *Cell) {
		row := c.row
		col := c.column
		c.north = g.cellAt(row-1, col)
		c.south = g.cellAt(row+1, col)
		c.west = g.cellAt(row, col-1)
		c.east = g.cellAt(row, col+1)
	})
	return
}

func (g *Grid) cellAt(row, column int) *Cell {
	if row >= 0 && row < g.rows {
		if column >= 0 && column < len(g.grid[row]) {
			return g.grid[row][column]
		}
	}
	return nil
}

func (g *Grid) randomCell() *Cell {
	row := rand.Intn(g.rows)
	column := rand.Intn(len(g.grid[row]))
	return g.cellAt(row, column)
}

func (g Grid) size() int {
	return g.rows * g.columns
}

func (g *Grid) eachRow(fn func([]*Cell)) {
	for _, row := range g.grid {
		fn(row)
	}
}

func (g *Grid) eachCell(fn func(*Cell)) {
	g.eachRow(func(row []*Cell) {
		for _, sell := range row {
			fn(sell)
		}
	})
}

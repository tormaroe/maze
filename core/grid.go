package core

import (
	"math/rand"
)

type Grid struct {
	Rows      int
	Columns   int
	grid      [][]*Cell
	Distances *Distance
}

func NewGrid(rows, columns int) (g Grid) {
	g = Grid{Rows: rows, Columns: columns}
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
		c.north = g.CellAt(row-1, col)
		c.south = g.CellAt(row+1, col)
		c.west = g.CellAt(row, col-1)
		c.east = g.CellAt(row, col+1)
	})
	return
}

func (g *Grid) CellAt(row, column int) *Cell {
	if row >= 0 && row < g.Rows {
		if column >= 0 && column < len(g.grid[row]) {
			return g.grid[row][column]
		}
	}
	return nil
}

func (g *Grid) randomCell() *Cell {
	row := rand.Intn(g.Rows)
	column := rand.Intn(len(g.grid[row]))
	return g.CellAt(row, column)
}

func (g Grid) size() int {
	return g.Rows * g.Columns
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

func (g *Grid) DeadendsCount() int {
	count := 0
	g.eachCell(func(c *Cell) {
		if len(c.links) == 1 {
			count++
		}
	})
	return count
}

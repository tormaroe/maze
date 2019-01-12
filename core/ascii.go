package core

import (
	"strconv"
	"strings"
)

func (g Grid) contentsOfCell(c *Cell) string {
	if g.Distances != nil {
		if v, found := g.Distances.distanceTo(c); found {
			return strconv.FormatInt(int64(v), 36)
		}
	}
	return " "
}

func (g Grid) String() string {
	var sb strings.Builder
	sb.WriteString("+")
	for i := 0; i < g.columns; i++ {
		sb.WriteString("---+")
	}
	sb.WriteString("\n")

	g.eachRow(func(row []*Cell) {
		top := "|"
		bottom := "+"

		for _, c := range row {
			if c == nil {
				c = newCell(-1, -1)
			}

			body := " " + g.contentsOfCell(c) + " "
			eastBoundary := "|"
			if c.isLinked(c.east) {
				eastBoundary = " "
			}
			top += body + eastBoundary

			southBoundary := "---"
			if c.isLinked(c.south) {
				southBoundary = "   "
			}
			corder := "+"
			bottom += southBoundary + corder
		}
		sb.WriteString(top)
		sb.WriteString("\n")
		sb.WriteString(bottom)
		sb.WriteString("\n")
	})

	return sb.String()
}

package core

import "strings"

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

			body := "   "
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

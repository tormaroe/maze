package core

import "github.com/fogleman/gg"

func (g Grid) Png(path string) error {
	cellSize := 5
	imgWidth := cellSize * g.columns
	imgHeight := cellSize * g.rows
	dc := gg.NewContext(imgWidth+2, imgHeight+2)

	dc.SetRGB(237, 234, 232)
	dc.DrawRectangle(0.0, 0.0, float64(imgWidth), float64(imgHeight))
	dc.Fill()

	dc.SetLineWidth(1.0)
	dc.SetRGB255(102, 83, 69)
	dc.SetLineCapSquare()

	g.eachCell(func(c *Cell) {
		x1 := float64(c.column*cellSize + 1)
		y1 := float64(c.row*cellSize + 1)
		x2 := float64((c.column+1)*cellSize + 1)
		y2 := float64((c.row+1)*cellSize + 1)

		if c.north == nil {
			dc.DrawLine(x1, y1, x2, y1)
			dc.Stroke()
		}
		if c.west == nil {
			dc.DrawLine(x1, y1, x1, y2)
			dc.Stroke()
		}
		if !c.isLinked(c.east) {
			dc.DrawLine(x2, y1, x2, y2)
			dc.Stroke()
		}
		if !c.isLinked(c.south) {
			dc.DrawLine(x1, y2, x2, y2)
			dc.Stroke()
		}
	})

	return dc.SavePNG(path)
}

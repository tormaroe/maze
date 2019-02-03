package core

import (
	"image/color"
	"math"

	"github.com/fogleman/gg"
)

const CellSize = 10
const WallThickness = 2.0

func (g Grid) backgroundColor(c *Cell) (int, int, int) {
	distance, found := g.Distances.distanceTo(c)
	if found {
		_, maximum := g.Distances.Max()
		intensity := float64(maximum-distance) / float64(maximum)
		dark := math.Round(255 * intensity)
		bright := 128 + math.Round(127*intensity)
		return int(dark), int(bright), int(dark)
	}
	return 0, 0, 0
}

func (g Grid) PngNewContext() *gg.Context {
	imgWidth := CellSize * g.Columns
	imgHeight := CellSize * g.Rows
	dc := gg.NewContext(imgWidth+2, imgHeight+2)
	return dc
}

func SetBackground(ctx *gg.Context, r, g, b float64) {
	ctx.SetRGB(r, g, b)
	ctx.DrawRectangle(0.0, 0.0, float64(ctx.Width()), float64(ctx.Height()))
	ctx.Fill()
}

func (g Grid) Png(path string) error {
	cellSize := 5
	imgWidth := cellSize * g.Columns
	imgHeight := cellSize * g.Rows
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

func (g Grid) PngDrawWalls(dc *gg.Context) {
	dc.SetLineWidth(WallThickness)
	dc.SetColor(color.Black)

	g.eachCell(func(c *Cell) {
		x1 := float64(c.column*CellSize + 1)
		y1 := float64(c.row*CellSize + 1)
		x2 := float64((c.column+1)*CellSize + 1)
		y2 := float64((c.row+1)*CellSize + 1)

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

	return
}

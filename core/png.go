package core

import (
	"image/color"
	"math"

	"github.com/fogleman/gg"
)

const CellSize = 10
const WallThickness = 2.0

func (g Grid) backgroundColor(c *Cell) (float64, float64, float64) {
	distance, found := g.Distances.distanceTo(c)
	if found {
		_, maximum := g.Distances.Max()
		intensity := float64(maximum-distance) / float64(maximum)
		dark := 255 - math.Round(254*intensity)
		bright := 255 - math.Round(90*intensity)
		return bright, dark, bright
	}
	return 255, 255, 255
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

func (c Cell) dimensions() (x1, y1, x2, y2 float64) {
	x1 = float64(c.column*CellSize + 1)
	y1 = float64(c.row*CellSize + 1)
	x2 = float64((c.column+1)*CellSize + 1)
	y2 = float64((c.row+1)*CellSize + 1)
	return
}

func (g Grid) PngDrawWalls(dc *gg.Context) {
	dc.SetLineWidth(WallThickness)
	dc.SetColor(color.Black)

	g.eachCell(func(c *Cell) {
		x1, y1, x2, y2 := c.dimensions()

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

func (g Grid) PngColorizeCells(dc *gg.Context) {
	g.eachCell(func(c *Cell) {
		x1, y1, x2, y2 := c.dimensions()
		r, g, b := g.backgroundColor(c)
		dc.SetRGB(r, g, b)
		dc.DrawRectangle(x1, y1, x2-x1, y2-y1)
		dc.Fill()
	})
}

func (g Grid) PngDotPath(dc *gg.Context) {
	if g.Distances == nil {
		return
	}
	dc.SetRGB(10, 250, 250)
	g.eachCell(func(c *Cell) {
		if _, found := g.Distances.distanceTo(c); found {
			x1, y1, x2, y2 := c.dimensions()
			centerX := x2 - ((x2 - x1) / 2)
			centerY := y2 - ((y2 - y1) / 2)
			dc.DrawCircle(centerX, centerY, 2.5)
			dc.Fill()

		}
	})
}

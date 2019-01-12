package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/tormaroe/maze/core"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("\nBinary Tree Maze:")
	grid := core.NewGrid(5, 8)
	grid.BinaryTreeMaze()
	fmt.Print(grid)

	fmt.Println("\nSidewinder Maze:")
	grid = core.NewGrid(5, 8)
	grid.SidewinderMaze()
	fmt.Print(grid)

	fmt.Println("\nBinary Tree Maze with distances:")
	rand.Seed(15302)
	dg := core.NewGrid(6, 10)
	dg.SidewinderMaze()
	start := dg.CellAt(0, 0)
	dg.Distances = start.Distances()
	fmt.Print(dg)

	fmt.Println("\nPath from northwest corner to southwest corner:")
	rand.Seed(15302)
	dg = core.NewGrid(6, 10)
	dg.SidewinderMaze()
	start = dg.CellAt(0, 0)
	d := start.Distances()
	dg.Distances = d.PathToCell(dg.CellAt(dg.Rows-1, 0))
	fmt.Print(dg)

	fmt.Println("\nLongest path:")
	rand.Seed(15302)
	dg = core.NewGrid(6, 10)
	dg.SidewinderMaze()
	start = dg.CellAt(0, 0)
	distances := start.Distances()
	newStart, _ := distances.Max()
	newDistances := newStart.Distances()
	goal, _ := newDistances.Max()
	dg.Distances = newDistances.PathToCell(goal)
	fmt.Print(dg)

	path := "out" + strconv.FormatInt(rand.Int63(), 10) + ".png"
	fmt.Println("Saving 20x30 Sidewinder maze to " + path)
	grid = core.NewGrid(20, 30)
	grid.SidewinderMaze()
	if err := grid.Png(path); err != nil {
		panic(err)
	}
}

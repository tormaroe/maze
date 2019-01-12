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
	dg := core.NewGrid(6, 10)
	dg.SidewinderMaze()
	start := dg.CellAt(0, 0)
	dg.Distances = start.Distances()
	fmt.Print(dg)

	path := "out" + strconv.FormatInt(rand.Int63(), 10) + ".png"
	fmt.Println("Saving 20x30 Sidewinder maze to " + path)
	grid = core.NewGrid(20, 30)
	grid.SidewinderMaze()
	if err := grid.Png(path); err != nil {
		panic(err)
	}
}

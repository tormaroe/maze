package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/docopt/docopt-go"
	"github.com/tormaroe/maze/core"
)

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func main() {

	usage := `Maze generation tool inspired by the book Mazes for Programmers.
	
Usage:
  maze algorithms
  maze ascii <algorithm> [--width <n>] [--height <n>] [--seed <n>] [-vrl]
  maze png <algorithm> [--out <path>] [--width <n>] [--height <n>] [--seed <n>] [-vrlc]
  maze --help

Options:
  -v, --verbous                  Print verbous information.
  -w <n>, --width <n>            Maze width (number of columns) [default: 10].
  -H <n>, --height <n>           Maze height (number of rows) [default: 8].
  -s <n>, --seed <n>             Random seed [default: 0].
  -r, --random                   Use a random random seed, no matter what seed has been specified.
  -l, --longest-path             Calculate and mark longest path.
  -c, --colorize                 Colorize cells based on distance.
  -o <path>, --out <path>        [default: out.png]`

	arguments, _ := docopt.ParseDoc(usage)

	verbous, _ := arguments.Bool("--verbous")

	if verbous {
		fmt.Print("OPTIONS: ")
		fmt.Println(prettyPrint(arguments))
	}

	if listAlgos, _ := arguments.Bool("algorithms"); listAlgos {
		fmt.Println("Avaliable algorithms are:")
		fmt.Println("  * binarytree")
		fmt.Println("  * sidewinder")
		fmt.Println("  * aldousbroder")
		return
	}

	seed := int64(0)
	if random, _ := arguments.Bool("--random"); random {
		seed = time.Now().UnixNano()
	} else {
		seedTemp, ok := arguments.Int("--seed")
		if ok != nil {
			panic("Invalid seed value.")
		}
		seed = int64(seedTemp)
	}

	if verbous {
		fmt.Printf("Random seed: %d\n", seed)
	}
	rand.Seed(seed)

	ascii, _ := arguments.Bool("ascii")
	png, _ := arguments.Bool("png")
	width, _ := arguments.Int("--width")
	height, _ := arguments.Int("--height")
	algorithm, _ := arguments.String("<algorithm>")
	longestPath, _ := arguments.Bool("--longest-path")
	colorize, _ := arguments.Bool("--colorize")

	grid := core.NewGrid(height, width)

	switch algorithm {
	case "sidewinder":
		grid.SidewinderMaze()
		break
	case "binarytree":
		grid.BinaryTreeMaze()
		break
	case "aldousbroder":
		grid.AldousBroderMaze()
		break
	}

	if longestPath {
		grid.Distances = grid.LongestPath()
	}

	if ascii {
		fmt.Print(grid)
	}

	if png {
		path, _ := arguments.String("--out")
		dc := grid.PngNewContext()
		core.SetBackground(dc, 20, 20, 20)

		if colorize {
			start := grid.CellAt(grid.Rows/2, grid.Columns/2)
			grid.Distances = start.Distances()
			grid.PngColorizeCells(dc)
		}

		grid.PngDrawWalls(dc)

		if longestPath {
			grid.Distances = grid.LongestPath()
			grid.PngDotPath(dc)
		}

		if verbous {
			fmt.Println("Writing", path)
		}
		dc.SavePNG(path)
	}
}

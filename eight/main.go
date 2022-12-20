package eight

import (
	"strings"
)

type Coordinate struct {
	x, y int
}

func makeRange(min, max int) []int { // Generates an inclusive range
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func calculateVisibility(tg [][]string, xr []int, yr []int, x int, y int) bool {
	for _, j := range yr {
		for _, i := range xr {
			if tg[i][j] >= tg[x][y] {
				return false
			}
		}
	}

	return true
}

func main(file []byte) int {
	treesSource := string(file)
	treeGrid := strings.Split(treesSource, "\n")

	w := len(treeGrid[0])
	h := len(treeGrid)
	v := []Coordinate{}

	tg := [][]string{}
	for _, row := range treeGrid {
		tg = append(tg, strings.Split(row, ""))
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if y == 0 || y == h-1 || x == 0 || x == w-1 {
				v = append(v, Coordinate{x, y})
				continue
			}

			visible := calculateVisibility(tg, makeRange(0, x-1), []int{y}, x, y) // Left

			if visible {
				v = append(v, Coordinate{x, y})
				continue
			}

			visible = calculateVisibility(tg, makeRange(x+1, w-1), []int{y}, x, y) // Right

			if visible {
				v = append(v, Coordinate{x, y})
				continue
			}

			visible = calculateVisibility(tg, []int{x}, makeRange(0, y-1), x, y) // Top

			if visible {
				v = append(v, Coordinate{x, y})
				continue
			}

			visible = calculateVisibility(tg, []int{x}, makeRange(y+1, h-1), x, y) // Bottom

			if visible {
				v = append(v, Coordinate{x, y})
			}
		}
	}

	return len(v)
}

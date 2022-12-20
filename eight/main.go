package eight

import (
	"strings"
)

type Coordinate struct {
	x, y int
}

func makeRange(min, max int, rev bool) []int {
	a := make([]int, max-min)
	for i := range a {
		if rev {
			a[i] = max - i
		} else {
			a[i] = min + i
		}
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

			lv := calculateVisibility(tg, makeRange(0, x, false), []int{y}, x, y)

			if lv {
				v = append(v, Coordinate{x, y})
				continue
			}

			rv := calculateVisibility(tg, makeRange(x, w-1, true), []int{y}, x, y)

			if rv {
				v = append(v, Coordinate{x, y})
				continue
			}

			tv := calculateVisibility(tg, []int{x}, makeRange(0, y, false), x, y)

			if tv {
				v = append(v, Coordinate{x, y})
				continue
			}

			bv := calculateVisibility(tg, []int{x}, makeRange(y, h-1, true), x, y)

			if bv {
				v = append(v, Coordinate{x, y})
			}
		}
	}

	return len(v)
}

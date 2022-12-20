package eight

import (
	"strings"
)

type Coordinate struct {
	x, y int
}

func main(file []byte) int {
	treesSource := string(file)
	treeGrid := strings.Split(treesSource, "\n")

	w := len(treeGrid[0])
	h := len(treeGrid)
	v := []Coordinate{}

	treesSource = strings.ReplaceAll(treesSource, "\n", "")
	treeGrid = strings.Split(treesSource, "")

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if y == 0 || y == h-1 || x == 0 || x == w-1 {
				v = append(v, Coordinate{x, y})
				continue
			}

			lv, rv, tv, bv := true, true, true, true

			// left check
			for cx := y * w; cx < (y*w)+x; cx++ {
				if treeGrid[cx] >= treeGrid[y*w+x] {
					lv = false
					break
				}
			}

			if lv {
				v = append(v, Coordinate{x, y})
				continue
			}

			// right check
			for cx := (y+1)*w - 1; cx > y*w+x; cx-- {
				if treeGrid[cx] >= treeGrid[y*w+x] {
					rv = false
					break
				}
			}

			if rv {
				v = append(v, Coordinate{x, y})
				continue
			}

			// top check
			for cy := x; cy < x+(y*w); cy += w {
				if treeGrid[cy] >= treeGrid[y*w+x] {
					tv = false
					break
				}
			}

			if tv {
				v = append(v, Coordinate{x, y})
				continue
			}

			// bottom check
			for cy := x + ((h - 1) * w); cy > x; cy -= w {
				if treeGrid[cy] >= treeGrid[y*w+x] {
					bv = false
					break
				}
			}

			if bv {
				v = append(v, Coordinate{x, y})
				continue
			}
		}
	}

	return len(v)
}

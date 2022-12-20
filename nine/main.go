package nine

import (
	"fmt"
	"strconv"
	"strings"
)

type moveset struct {
	dir  string
	dist int
}

type coordinate struct {
	x, y int
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func formatMoveSets(moves []string) []moveset {
	movesets := []moveset{}

	for _, move := range moves {
		parts := strings.Split(move, " ")
		direction := parts[0]
		travel, err := strconv.Atoi(parts[1])
		check(err)

		movesets = append(movesets, moveset{dir: direction, dist: travel})
	}

	return movesets
}

func intabs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func printRoute(wl int, wh int, hl int, hh int, tpos map[coordinate]int) {
	route := [][]string{}
	w := wh - wl
	h := hh - hl

	for i := 0; i <= h; i++ {
		row := []string{}
		for j := 0; j <= w; j++ {
			row = append(row, ".")
		}
		route = append(route, row)
	}

	keys := make([]coordinate, 0, len(tpos))
	for c := range tpos {
		keys = append(keys, c)
	}

	for _, key := range keys {
		offsety := h - (key.y + intabs(hl))
		offsetx := key.x + intabs(wl)
		route[offsety][offsetx] = "#"
	}

	fmt.Println("\nTail route:")
	for _, routerow := range route {
		fmt.Println(routerow)
	}
}

func adjustDimensions(wl int, wh int, hl int, hh int, wt int, ht int) (int, int, int, int) {
	if wt > wh {
		wh = wt
	} else if wt < wl {
		wl = wt
	}

	if ht > hh {
		hh = ht
	} else if ht < wl {
		hl = ht
	}

	return wl, wh, hl, hh
}

func main(file []byte) int {
	moves := strings.Split(string(file), "\n")

	h := coordinate{0, 0}
	t := coordinate{0, 0}
	tpos := make(map[coordinate]int)
	movesets := formatMoveSets(moves)
	wl, wh, wt, hl, hh, ht := 0, 0, 0, 0, 0, 0

	for _, moveset := range movesets {
		for i := 0; i < moveset.dist; i++ {
			switch moveset.dir {
			case "U":
				h = coordinate{h.x, h.y + 1}
				ht++
			case "D":
				h = coordinate{h.x, h.y - 1}
				ht--
			case "L":
				h = coordinate{h.x - 1, h.y}
				wt--
			case "R":
				h = coordinate{h.x + 1, h.y}
				wt++
			}

			wl, wh, hl, hh = adjustDimensions(wl, wh, hl, hh, wt, ht)

			xdist := h.x - t.x
			ydist := h.y - t.y

			if xdist > 1 {
				t = coordinate{t.x + 1, t.y + ydist}
			} else if xdist < -1 {
				t = coordinate{t.x - 1, t.y + ydist}
			} else if ydist > 1 {
				t = coordinate{t.x + xdist, t.y + 1}
			} else if ydist < -1 {
				t = coordinate{t.x + xdist, t.y - 1}
			}

			tpos[t] += 1
		}
	}

	printRoute(wl, wh, hl, hh, tpos)

	return len(tpos)
}

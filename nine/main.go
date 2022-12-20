package nine

import (
	"strconv"
	"strings"
)

type coordinate struct {
	x, y int
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main(file []byte) int {
	moves := strings.Split(string(file), "\n")

	h := coordinate{0, 0}
	t := coordinate{0, 0}
	tpos := make(map[coordinate]int)

	for _, move := range moves {
		parts := strings.Split(move, " ")
		direction := parts[0]
		travel, err := strconv.Atoi(parts[1])
		check(err)

		for i := 0; i < travel; i++ {
			switch direction {
			case "U":
				h = coordinate{h.x, h.y + 1}
			case "D":
				h = coordinate{h.x, h.y - 1}
			case "L":
				h = coordinate{h.x - 1, h.y}
			case "R":
				h = coordinate{h.x + 1, h.y}
			}

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

	return len(tpos)
}

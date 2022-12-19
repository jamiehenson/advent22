package main

import (
	"fmt"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	file, err := os.ReadFile("src/signals.txt")
	check(err)

	signals := strings.Split(string(file), "\n")

	for si, signal := range signals { // Each signal line from file
		if len(signal) <= 4 {
			continue
		}

		marker_found, marker_index := false, 0

		for i := 0; i < len(signal)-4; i++ { // Each window from 0 to n-4
			char_map := make(map[byte]int)
			distinct := true

			for j := 0; j < 4; j++ { // For each character in the four char slice
				char_map[signal[i+j]] += 1
			}

			for _, v := range char_map {
				if v >= 2 {
					distinct = false
				}
			}

			if distinct {
				marker_found = true
				marker_index = i + 4
				break
			}
		}

		if marker_found {
			fmt.Printf("Signal %d: first marker is at character %d\n", si, marker_index)
		} else {
			fmt.Printf("No markers detected in Signal %d\n", si)
		}
	}
}

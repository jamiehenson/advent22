package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func makeRange(l int, u int) []int {
	r := make([]int, u-l+1)
	for i := range r {
		r[i] = i + l
	}

	return r
}

func main() {
	file, err := os.ReadFile("src/cleaning.txt")
	check(err)

	pairs := strings.Split(string(file), "\n")
	pair_matches := []bool{}
	pair_match_count := 0

	for _, pair := range pairs {
		elves := strings.Split(pair, ",")

		ranges := [][]int{}
		match_count := 0

		for _, elf := range elves {
			bounds := strings.Split(elf, "-")
			lower_bound, err := strconv.Atoi(bounds[0])
			check(err)
			upper_bound, err := strconv.Atoi(bounds[1])
			check(err)

			ranges = append(ranges, makeRange(lower_bound, upper_bound))
		}

		sort.Slice(ranges, func(i, j int) bool {
			return len(ranges[i]) > len(ranges[j])
		})

		lm := make(map[int]bool)

		for _, i := range ranges[0] {
			lm[i] = true
		}

		for _, j := range ranges[1] {
			if lm[j] {
				match_count += 1
			}
		}

		match := match_count == len(ranges[1])
		pair_matches = append(pair_matches, match)

		if match {
			pair_match_count += 1
		}
	}

	fmt.Println(pair_matches, pair_match_count)
}

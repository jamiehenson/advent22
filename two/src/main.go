package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type round struct {
	me   int
	them int
}

func calculateResult(r round) (points int, err error) {
	// Rock 0
	// Paper 1
	// Scissors 2

	if r.me == r.them { // Draw
		points = 3
	} else if (r.me-1)%3 == r.them { // Win
		points = 6
	} else if (r.me+1)%3 == r.them { // Loss
		points = 0
	} else {
		err = errors.New("unexpected result condition")
	}

	return
}

func encodeMove(m byte) (val int, err error) {
	switch m {
	case []byte("X")[0], []byte("A")[0]: // Rock
		val = 0
	case []byte("Y")[0], []byte("B")[0]: // Paper
		val = 1
	case []byte("Z")[0], []byte("C")[0]: // Scissors
		val = 2
	default:
		err = errors.New("incorrect string value")
	}

	return
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func constructRound(row string) round {
	m := row[2]
	t := row[0]

	em, err := encodeMove(m)
	check(err)

	et, err := encodeMove(t)
	check(err)

	r := round{me: em, them: et} // e.g. {me: 0, them: 1}

	return r
}

func sum(results []int) int {
	subtotal := 0

	for _, result := range results {
		subtotal += result
	}

	return subtotal
}

func main() {
	file, err := os.ReadFile("src/scores.txt")
	check(err)

	rows := strings.Split(string(file), "\n")
	results := []int{}

	for index, row := range rows {
		r := constructRound(row)

		move_value := r.me + 1

		result, err := calculateResult(r)
		check(err)

		subtotal := move_value + result

		results = append(results, subtotal)
		fmt.Printf("Round %d: %d\n", index+1, subtotal)
	}

	fmt.Printf("Total score: %d\n", sum(results))
}

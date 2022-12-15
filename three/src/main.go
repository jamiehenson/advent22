package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func calculatePriority(item rune) (p int, err error) {
	a := int(item)

	if a >= 97 && a <= 122 { // lowercase
		p = a - 96
	} else if a >= 65 && a <= 90 { // uppercase
		p = a - 38
	} else {
		err = errors.New("incompatible char supplied")
	}

	return
}

func sliceContains(s []rune, v rune) bool {
	for _, r := range s {
		if r == v {
			return true
		}
	}

	return false
}

func calculateIntersection(a string, b string) (i []rune, err error) {
	h := make(map[rune]int, len(a))

	for _, c := range a {
		h[c] += 1
	}

	for _, c := range b {
		if h[c] > 0 && !sliceContains(i, c) {
			i = append(i, c)
		}
	}

	return
}

func main() {
	file, err := os.ReadFile("src/backpacks.txt")
	check(err)

	backpacks := strings.Split(string(file), "\n")
	ti := [][]rune{}

	for _, b := range backpacks {
		m := len(b) / 2
		s := b[:m]
		e := b[m:]

		i, err := calculateIntersection(s, e)
		ti = append(ti, i)
		check(err)
	}

	// Flattened vals
	fi := []string{}
	pt := 0

	for _, rs := range ti {
		for _, r := range rs {
			s := string(r)
			fi = append(fi, s)
			p, err := calculatePriority(r)
			check(err)

			pt += p
		}
	}

	fmt.Printf("Common items: %v\n", fi)
	fmt.Printf("Priority total: %d\n", pt)
}

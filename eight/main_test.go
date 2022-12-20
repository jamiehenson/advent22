package eight

import (
	"os"
	"testing"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func TestMain(t *testing.T) {
	file, err := os.ReadFile("src/trees.txt")
	check(err)

	ans := main(file)
	if ans != 21 {
		t.Errorf("%d visible trees detected; want 21", ans)
	}
}

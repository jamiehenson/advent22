package nine

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	file, err := os.ReadFile("moves.txt")
	check(err)

	ans := main(file)

	if ans != 13 {
		t.Errorf("Got %d. Expected 13.", ans)
	}
}

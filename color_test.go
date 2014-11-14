package color

import (
	"testing"
)

func Example() {
	Println(Red, "Red", Green, " Green")
	Println("Part1", Blue, " Part", "2", None, " OK ", Red, "!")
}

func TestPrintln(t *testing.T) {
	Example()
}

package color

import (
	"testing"
)

func Example() {
	Println(Red("Red "), Green("Green"))
	Println(Magenta("Normal"), Blue(" %s", "Blue"))
	Println("Normal", Blue(" %s", "Blue"))
}

func TestPrintln(t *testing.T) {
	Example()
}

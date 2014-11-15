package color

import (
	"testing"
)

func Example() {
	Println(Red("Red "), Green("Green"))
	Println(None("Normal"), Blue(" %s", "Blue"))
}

func TestPrintln(t *testing.T) {
	Example()
}

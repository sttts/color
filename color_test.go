package color

import (
	"testing"
)

func Example() {
	Println(Red("Red"), None(" %s ", "Normal"), Green("Green"))
}

func TestPrintln(t *testing.T) {
	Example()
}

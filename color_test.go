package color

import (
	"testing"
)

func Example() {
	Println(
		Green("Green"), " ",
		Red("Red").BgWhite(), " ",
		"Normal ",
		Yellow("Yellow"), " ",
		Blue("%s", "Blue"))
}

func TestForeAndBack(t *testing.T) {
	Println(Red("Red on White").BgWhite())
}

func TestForeground(t *testing.T) {
	Println(
		None("None"), " ",
		Black("Black"), " ",
		Red("Red"), " ",
		Green("Green"), " ",
		Yellow("Yellow"), " ",
		Blue("Blue"), " ",
		Magenta("Magenta"), " ",
		Cyan("Cyan"), " ",
		White("White"))

}

func TestBackgroud(t *testing.T) {
	Println(
		None("None").BgNone(), " ",
		None("Black").BgBlack(), " ",
		None("Red").BgRed(), " ",
		None("Green").BgGreen(), " ",
		None("Yellow").BgYellow(), " ",
		None("Blue").BgBlue(), " ",
		None("Magenta").BgMagenta(), " ",
		None("Cyan").BgCyan(), " ",
		None("White").BgWhite())
}

func TestNoColor(t *testing.T) {
	Println("No color")
}

func TestFormatColor(t *testing.T) {
	Println(Blue("%s", "Blue text with format"))
}

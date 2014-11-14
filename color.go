//Color is a simple package for printing in color to a windows or ansi console.
//Internally the package github.com/daviddengcn/go-colortext is used.
package color

import (
	"fmt"
	ct "github.com/daviddengcn/go-colortext"
)

//Possible colors
const (
	None = ct.Color(iota)
	Black
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

//Println prints text to terminal with colors.
//The parameters are string or colors in any order.
//At the end of the line the color will be reset.
func Println(v ...interface{}) {
	for _, param := range v {
		switch v := param.(type) {
		default:
			fmt.Printf("unexpected type %T", v)
		case string:
			fmt.Print(param)
			break
		case ct.Color:
			color := param.(ct.Color)
			if color == None {
				ct.ResetColor()
			} else {
				ct.ChangeColor(color, false, ct.None, false)
			}
		}
	}
	ct.ResetColor()
	fmt.Println()
}

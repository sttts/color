//Color is a simple package for printing in color to a windows or ansi console.
//Internally the package github.com/daviddengcn/go-colortext is used.
package color

import (
	"fmt"
	ct "github.com/daviddengcn/go-colortext"
)

//Println prints text to terminal with colors.
//At the end of the line the color will be reset.
func Println(msgs ...ColorMsg) {
	for _, colorMsg := range msgs {

		if colorMsg.Color == ct.None {
			ct.ResetColor()
		} else {
			ct.ChangeColor(colorMsg.Color, false, ct.None, false)
		}

		fmt.Print(colorMsg.Message)
	}
	ct.ResetColor()
	fmt.Println()
}

//ColorMsg contains the message and color to use
type ColorMsg struct {
	Color   ct.Color
	Message string
}

func (colorMsg ColorMsg) String() string {
	return colorMsg.Message
}

func newColorMsg(color ct.Color, format string, v ...interface{}) (msg ColorMsg) {
	msg.Color = color
	msg.Message = fmt.Sprintf(format, v...)
	return
}

func None(format string, v ...interface{}) (msg ColorMsg) {
	return newColorMsg(ct.None, format, v...)
}

func Black(format string, v ...interface{}) (msg ColorMsg) {
	return newColorMsg(ct.Black, format, v...)
}

func Red(format string, v ...interface{}) (msg ColorMsg) {
	return newColorMsg(ct.Red, format, v...)
}

func Green(format string, v ...interface{}) (msg ColorMsg) {
	return newColorMsg(ct.Green, format, v...)
}

func Yellow(format string, v ...interface{}) (msg ColorMsg) {
	return newColorMsg(ct.Yellow, format, v...)
}

func Blue(format string, v ...interface{}) (msg ColorMsg) {
	return newColorMsg(ct.Blue, format, v...)
}

func Magenta(format string, v ...interface{}) (msg ColorMsg) {
	return newColorMsg(ct.Magenta, format, v...)
}

func Cyan(format string, v ...interface{}) (msg ColorMsg) {
	return newColorMsg(ct.Cyan, format, v...)
}

func White(format string, v ...interface{}) (msg ColorMsg) {
	return newColorMsg(ct.White, format, v...)
}

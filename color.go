//Color is a simple package for printing in color to a windows or ansi console.
//Internally the package github.com/daviddengcn/go-colortext is used.
package color

import (
	"fmt"
	ct "github.com/daviddengcn/go-colortext"
)

//Println prints text to terminal with colors.
//At the end of the line the color will be reset.
func Println(msgs ...interface{}) {
	for _, msg := range msgs {

		color := ct.None
		bgcolor := ct.None
		message := ""

		switch msg.(type) {
		case ColorMsg:
			colorMsg := msg.(ColorMsg)
			color = colorMsg.Color
			bgcolor = colorMsg.BgColor
			message = colorMsg.Message
			break
		default:
			message = fmt.Sprintf("%v", msg)
		}

		if color == ct.None && bgcolor == ct.None {
			ct.ResetColor()
		} else {
			ct.ChangeColor(color, false, bgcolor, false)
		}

		fmt.Print(message)
	}
	ct.ResetColor()
	fmt.Println()
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

//ColorMsg contains the message and color to use
type ColorMsg struct {
	Color   ct.Color
	BgColor ct.Color
	Message string
}

func newColorMsg(color ct.Color, format string, v ...interface{}) (msg ColorMsg) {
	msg.Color = color
	msg.BgColor = ct.None
	msg.Message = fmt.Sprintf(format, v...)
	return
}

func (colorMsg ColorMsg) String() string {
	return colorMsg.Message
}

//BgNone is not necessary since that is the default
func (colorMsg ColorMsg) BgNone() ColorMsg {
	colorMsg.BgColor = ct.None
	return colorMsg
}

func (colorMsg ColorMsg) BgBlack() ColorMsg {
	colorMsg.BgColor = ct.Black
	return colorMsg
}

func (colorMsg ColorMsg) BgRed() ColorMsg {
	colorMsg.BgColor = ct.Red
	return colorMsg
}
func (colorMsg ColorMsg) BgGreen() ColorMsg {
	colorMsg.BgColor = ct.Green
	return colorMsg
}
func (colorMsg ColorMsg) BgYellow() ColorMsg {
	colorMsg.BgColor = ct.Yellow
	return colorMsg
}
func (colorMsg ColorMsg) BgBlue() ColorMsg {
	colorMsg.BgColor = ct.Blue
	return colorMsg
}
func (colorMsg ColorMsg) BgMagenta() ColorMsg {
	colorMsg.BgColor = ct.Magenta
	return colorMsg
}
func (colorMsg ColorMsg) BgCyan() ColorMsg {
	colorMsg.BgColor = ct.Cyan
	return colorMsg
}
func (colorMsg ColorMsg) BgWhite() ColorMsg {
	colorMsg.BgColor = ct.White
	return colorMsg
}

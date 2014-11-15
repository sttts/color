Color is a simple package for printing in color to a windows or ansi console.  
Internally the package github.com/daviddengcn/go-colortext is used.

## Usage

```Go
color.Println(color.Red("Red "), color.Green("Green"))
color.Println(color.Magenta("Normal"), color.Blue(" %s", "Blue"))
color.Println("Normal", color.Blue(" %s", "Blue"))
```

[GoDoc](https://godoc.org/github.com/inando/color)
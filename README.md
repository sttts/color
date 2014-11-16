Color is a simple package for printing in color to a windows or ansi console.  
Internally the package github.com/daviddengcn/go-colortext is used.

## Usage

```Go
	color.Println(
		color.Green("Green"), " ",
		color.Red("Red").BgWhite(), " ",
		"Normal ",
		color.Yellow("Yellow"), " ",
		color.Blue("%s", "Blue"))
```

[GoDoc](https://godoc.org/github.com/inando/color)
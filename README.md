Color is a simple package for printing in color to a windows or ansi console.  
Internally the package github.com/daviddengcn/go-colortext is used.

## Usage

```Go
color.Println(color.Red, "Red", color.Green, " Green")
color.Println("Part1", color.Blue, " Part", "2", color.None, " OK ", color.Red, "!")
```
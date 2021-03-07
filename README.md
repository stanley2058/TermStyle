# TermStyle
*Go module for constructing output style for Linux terminals.*

## Example
![example](example.gif)
```go
str := termstyle.Style(
    "Colored String",
    foreground.TrueColor(128, 50, 128),
    background.BrightYellow,
    effects.Italic,
    effects.Bold,
    effects.Blink,
    effects.CrossedOut,
)
fmt.Println(str)
```
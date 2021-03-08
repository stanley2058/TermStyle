package foreground

import "fmt"

// Foreground colors
const (
	Black   string = "30"
	Red     string = "31"
	Green   string = "32"
	Yellow  string = "33"
	Blue    string = "34"
	Magenta string = "35"
	Cyan    string = "36"
	White   string = "37"

	BrightBlack   string = "90"
	BrightRed     string = "91"
	BrightGreen   string = "92"
	BrightYellow  string = "93"
	BrightBlue    string = "94"
	BrightMagenta string = "95"
	BrightCyan    string = "96"
	BrightWhite   string = "97"
)

// TrueColor construct terminal usable true color string, not all terminal support this.
func TrueColor(r uint8, g uint8, b uint8) string {
	return fmt.Sprintf("38;%d;%d;%d", r, g, b)
}

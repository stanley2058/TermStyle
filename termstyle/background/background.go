package background

import "fmt"

// Background colors
const (
	Black   string = "40"
	Red     string = "41"
	Green   string = "42"
	Yellow  string = "43"
	Blue    string = "44"
	Magenta string = "45"
	Cyan    string = "46"
	White   string = "47"

	BrightBlack   string = "100"
	BrightRed     string = "101"
	BrightGreen   string = "102"
	BrightYellow  string = "103"
	BrightBlue    string = "104"
	BrightMagenta string = "105"
	BrightCyan    string = "106"
	BrightWhite   string = "107"
)

// TrueColorBackground construct terminal usable true color string for background
func TrueColorBackground(r uint8, g uint8, b uint8) string {
	return fmt.Sprintf("48;%d;%d;%d", r, g, b)
}

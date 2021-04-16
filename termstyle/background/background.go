package background

import (
	"fmt"

	"github.com/stanley2058/termstyle/utils"
)

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

// TrueColorBackground construct terminal usable true color string for background, not all terminal support this.
func TrueColorBackground(r uint8, g uint8, b uint8) string {
	return fmt.Sprintf("48;%d;%d;%d", r, g, b)
}

// HSL return converted true color string (0<= h <360, 0<= s,l <=1)
func HSL(h float64, s float64, l float64) string {
	return TrueColorBackground(utils.HslToRgb(h, s, l))
}

// Bg256 256 colors for background
//   0-7:      standard colors (Black, Red, Green, Yellow, Blue, Magenta, Cyan, White)
//   8-15:     high intensity colors
//   16-231:   16 + 36×r + 6×g + b (0 ≤ r, g, b ≤ 5)
//   232-255:  grayscale from black to white in 24 steps
//   see: "https://jonasjacek.github.io/colors"
func Bg256(num uint8) string {
	return fmt.Sprintf("\033[38;5;%dm", num)
}

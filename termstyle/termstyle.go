package termstyle

import "fmt"

const nc string = "\033[0m"

// Style style str with otps
func Style(str string, opts ...string) string {
	if len(opts) == 0 {
		return str
	}

	prefix := "\033[" + opts[0]
	for i := 1; i < len(opts); i++ {
		prefix += ";" + opts[i]
	}
	return fmt.Sprintf("%sm%s%s", prefix, str, nc)
}

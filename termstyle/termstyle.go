package termstyle

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

const nc string = "\033[0m"

var (
	initialized  = false
	supportColor = false
)

func init() {
	if initialized {
		return
	}

	initialized = true
	res, err := exec.Command("tput", "colors").Output()
	intRes, _ := strconv.Atoi(strings.TrimSpace(string(res)))

	if intRes > 2 && err == nil {
		supportColor = true
	}
}

// Style style str with opts
func Style(str string, opts ...string) string {
	if !supportColor || len(opts) == 0 {
		return str
	}

	prefix := "\033[" + opts[0]
	for i := 1; i < len(opts); i++ {
		prefix += ";" + opts[i]
	}
	return fmt.Sprintf("%sm%s%s", prefix, str, nc)
}

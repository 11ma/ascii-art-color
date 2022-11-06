package color

import (
	"os"
	"strings"
)

func IsFlagPassed(name string) bool {
	val := strings.Trim(os.Args[1], "-")
	flagName := strings.Split(val, "=")[0]

	return flagName == name
}

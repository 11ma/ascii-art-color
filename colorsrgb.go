package color

import (
	"os"
	"strings"
)

func RGBToANSI(r, g, b string) string {
	var color = "\x1b[38;2;" + r + ";" + g + ";" + b + "m"
	return color
}

func ColorRGB(s string) string {
	s = strings.ToLower(s)
	s = strings.TrimPrefix(s, "rgb(")
	s = strings.TrimSuffix(s, ")")
	removeSpaceAndJoin := strings.Join(strings.Fields(s), "")
	splits := strings.Split(removeSpaceAndJoin, ",")

	if len(splits) != 3 {
		os.Exit(0)
	}

	r := strings.Trim(splits[0], ",")
	g := strings.Trim(splits[1], ",")
	b := strings.Trim(splits[2], ",")

	return RGBToANSI(r, g, b)
}

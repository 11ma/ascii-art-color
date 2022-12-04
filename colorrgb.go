package color

import (
	"strconv"
	"strings"
)

// change rgb string value to ansi color
func RGBToANSI(r, g, b string) string {
	color := "\x1b[38;2;" + r + ";" + g + ";" + b + "m"
	return color
}

// get the rgb value in ansi code
func ColorRGB(s string) string {
	s = strings.ToLower(s)
	s = strings.TrimPrefix(s, "rgb(")
	s = strings.TrimSuffix(s, ")")
	removeSpaceAndJoin := strings.Join(strings.Fields(s), "")
	splits := strings.Split(removeSpaceAndJoin, ",")

	if len(splits) != 3 {
		WrongColorType()
	}

	for i := 0; i < len(splits); i++ {
		hexValue, err := strconv.Atoi(splits[i])
		if err != nil {
			WrongColorType()
		}
		if hexValue > 255 || hexValue < 0 {
			WrongColorType()
		}
	}

	r := strings.Trim(splits[0], ",")
	g := strings.Trim(splits[1], ",")
	b := strings.Trim(splits[2], ",")

	return RGBToANSI(r, g, b)
}

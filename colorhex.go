package color

import (
	"fmt"
	"strconv"
	"strings"
)

// helper function to convert anything to decimal
func anyToDecimal(value string, base int) string {
	decimal, err := strconv.ParseInt(value, base, 32)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	return strconv.Itoa(int(decimal))
}

// convert hex values to decimals
func hexToDecimal(hex string) string {
	return anyToDecimal(hex, 16)
}

// get the hex value in ansi code
func ColorHex(s string) string {
	s = strings.ToLower(s)
	s = strings.Trim(s, "#")
	if len(s) == 3 {
		newcolor := []string{s[0:1], s[0:1], s[1:2], s[1:2], s[2:], s[2:]}
		s = strings.Join(newcolor, "")
	}

	for _, digit := range s {
		if !((digit >= '0' && digit <= '9') || (digit >= 'a' && digit <= 'f')) {
			WrongColorType()
		}
	}

	if len(s) != 6 {
		WrongColorType()
	}

	r := s[0:2]
	g := s[2:4]
	b := s[4:]

	return RGBToANSI(hexToDecimal(r), hexToDecimal(g), hexToDecimal(b))
}

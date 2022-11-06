package color

import "strings"

func ColorHex(s string) string {
	s = strings.TrimPrefix(s, "#")
	if len(s) == 3 {
		newcolor := []string{s[0:1], s[0:1], s[1:2], s[1:2], s[2:], s[2:]}
		s = strings.Join(newcolor, "")
	}
	r := s[0:2]
	g := s[2:4]
	b := s[4:]
	return RGBToANSI(HexToDecimal(r), HexToDecimal(g), HexToDecimal(b))
}

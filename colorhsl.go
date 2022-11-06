package color

import (
	"math"
	"strconv"
	"strings"
)

func convertStringToFloat(s string) (H, S, L float64) {
	s = strings.TrimPrefix(s, "hsl(")
	s = strings.TrimSuffix(s, ")")
	removeSpaceAndJoin := strings.Join(strings.Fields(s), "")
	hsl := strings.Split(removeSpaceAndJoin, ",")

	hTrimmed, _ := strconv.Atoi(strings.Trim(hsl[0], "%,"))
	sTrimmed, _ := strconv.Atoi(strings.Trim(hsl[1], "%,"))
	lTrimmed, _ := strconv.Atoi(strings.Trim(hsl[2], "%,"))

	H = float64(hTrimmed)
	S = float64(sTrimmed) / 100
	L = float64(lTrimmed) / 100
	return H, S, L
}

func converToRGB(H, S, L float64) (C, X, m float64) {
	modC := math.Abs((2 * L) - 1)
	modX := math.Abs(math.Mod((H/60), 2) - 1)

	// we compute chroma, by multiplying saturation by the maximum chroma for a given lightness or value
	C = (1 - modC) * S
	// we find the point on one of the bottom three faces of the RGB cube which has the same hue and chroma as our color
	X = C * (1 - modX)

	m = L - C/2

	return C, X, m
}

func checkHueValue(H, C, X, m float64) (R, G, B string) {
	var R1 float64
	var G1 float64
	var B1 float64

	// checking the range of the hue value for lightness
	switch {
	case H < 60:
		R1 = C
		G1 = X
		B1 = 0
	case H < 120:
		R1 = X
		G1 = C
		B1 = 0
	case H < 180:
		R1 = 0
		G1 = C
		B1 = X
	case H < 240:
		R1 = 0
		G1 = X
		B1 = C
	case H < 300:
		R1 = X
		G1 = 0
		B1 = C
	case H < 360:
		R1 = C
		G1 = 0
		B1 = X
	}
	// we add equal amounts of R, G, and B to reach the proper lightness
	R = strconv.Itoa(int((R1 + m) * 255))
	G = strconv.Itoa(int((G1 + m) * 255))
	B = strconv.Itoa(int((B1 + m) * 255))

	return R, G, B
}

func ColorHSL(s string) string {
	H, S, L := convertStringToFloat(s)
	C, X, m := converToRGB(H, S, L)
	R, G, B := checkHueValue(H, C, X, m)
	return RGBToANSI(R, G, B)
}

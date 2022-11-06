package color

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func ColorHSL(s string) string {
	s = strings.TrimPrefix(s, "hsl(")
	s = strings.TrimSuffix(s, ")")
	hsl := strings.Fields(s)

	hTrimmed, _ := strconv.Atoi(strings.Trim(hsl[0], "%,"))
	sTrimmed, _ := strconv.Atoi(strings.Trim(hsl[1], "%,"))
	lTrimmed, _ := strconv.Atoi(strings.Trim(hsl[2], "%,"))

	H := float64(hTrimmed)
	S := float64(sTrimmed)
	L := float64(lTrimmed)

	S = S / 100
	L = L / 100

	modC := math.Abs((2 * L) - 1)
	modX := math.Abs(math.Mod((H/60), 2) - 1)

	C := (1 - modC) * S
	X := C * (1 - modX)
	m := L - C/2

	var RDerivative float64
	var GDerivative float64
	var BDerivative float64

	if H < 60 {
		RDerivative = C
		GDerivative = X
		BDerivative = 0
	} else if H < 120 {
		RDerivative = X
		GDerivative = C
		BDerivative = 0
	} else if H < 180 {
		RDerivative = 0
		GDerivative = C
		BDerivative = X
	} else if H < 240 {
		RDerivative = 0
		GDerivative = X
		BDerivative = C
	} else if H < 300 {
		RDerivative = X
		GDerivative = 0
		BDerivative = C
	} else if H < 360 {
		RDerivative = C
		GDerivative = 0
		BDerivative = X
	}

	R := int((RDerivative + m) * 255)
	G := int((GDerivative + m) * 255)
	B := int((BDerivative + m) * 255)

	fmt.Println(R, G, B)

	strR := strconv.Itoa(R)
	strG := strconv.Itoa(G)
	strB := strconv.Itoa(B)

	return RGBToANSI(strR, strG, strB)
}

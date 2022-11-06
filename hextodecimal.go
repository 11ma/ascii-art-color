package color

import (
	"fmt"
	"strconv"
)

func AnyToDecimal(value string, base int) string {
	decimal, err := strconv.ParseInt(value, base, 32)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	return strconv.Itoa(int(decimal))
}

// convert hex values to decimals
func HexToDecimal(hex string) string {
	return AnyToDecimal(hex, 16)
}

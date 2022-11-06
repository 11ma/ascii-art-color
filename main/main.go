package main

import (
	"color"
	"flag"
	"fmt"
)

func main() {
	word := "Hello"
	// target := len(word) - 4

	// for i := 0; i < len(word); i++ {
	// 	if i == target {
	// 		fmt.Print(color.Purple, string(word[i]), color.Reset)
	// 	} else {
	// 		fmt.Print(string(word[i]))
	// 	}
	// }
	// fmt.Println()

	colorFlag := flag.String("color", "", "set color")
	flag.Parse()

	// var colorAnsi = "\x1b[38;2;" + r + ";" + g + ";" + b + "m"

	// hsl
	// ansi := color.RGBToANSI(strR, strG, strB)
	// hslrgb := HslToRgb(rfloat, gfloat, bfloat)

	// var ansi = "\x1b[38;2;" + r + ";" + g + ";" + b + "m"
	fmt.Println(color.ColorHSL(*colorFlag), word)

}

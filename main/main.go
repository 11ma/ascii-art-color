package main

import (
	"color"
	"fmt"
)

func colorHex(r string) string {
	var color = "\x1b[38;2;" + r + "m"
	// fmt.Println(color)
	return color
}

func main() {
	word := "Hello"
	target := len(word) - 1

	for i := 0; i < len(word); i++ {
		if i == target {
			fmt.Print("\x1b[#ff0000m", string(word[i]), color.Reset)
		} else {
			fmt.Print(string(word[i]))
		}
	}
	fmt.Println()
}

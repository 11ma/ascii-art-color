package color

import (
	"fmt"
	"os"
)

func WrongColorType() {
	fmt.Println("Invalid color type")
	os.Exit(0)
}

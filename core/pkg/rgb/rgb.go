package rgb

import (
	"fmt"
	"math"
)

func createColor(i int) (int, int, int) {
	f := 0.1
	r := int(math.Sin(f*float64(i)+0)*127 + 128)
	g := int(math.Sin(f*float64(i)+2*math.Pi/3)*127 + 128)
	b := int(math.Sin(f*float64(i)+4*math.Pi/3)*127 + 128)
	return r, g, b
}

func Print(output []rune) {
	for char := 0; char < len(output); char++ {
		r, g, b := createColor(char)
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, output[char])
	}
	fmt.Println()
}

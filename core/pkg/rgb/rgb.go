package rgb

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

type Color struct {
	Red   int
	Green int
	Blue  int
}

func createColor(i int) *Color {
	f := 0.1
	return &Color{
		Red:   int(math.Sin(f*float64(i)+0)*127 + 128),
		Green: int(math.Sin(f*float64(i)+2*math.Pi/3)*127 + 128),
		Blue:  int(math.Sin(f*float64(i)+4*math.Pi/3)*127 + 128),
	}
}

func print(output []rune) {
	for j := 0; j < len(output); j++ {
		color := createColor(j)
		r, g, b := color.Red, color.Green, color.Blue
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, output[j])
	}
	fmt.Println()
}

func main() {
	info, _ := os.Stdin.Stat()
	var output []rune

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: fortune | gorainbow")
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	print(output)
}

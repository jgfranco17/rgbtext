package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	rgb "github.com/jgfranco17/rgbtext/core/pkg/rgb"
)

func readCliInput() []rune {
	var output []rune
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}
	return output
}

func main() {
	info, _ := os.Stdin.Stat()

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: echo 'Hello' | gorainbow")
		os.Exit(1)
	}

	output := readCliInput()
	rgb.Print(output)
}

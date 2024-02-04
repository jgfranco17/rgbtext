package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	rgb "github.com/jgfranco17/rgbtext/core/pkg/rgb"
)

func readCliInput() ([]rune, error) {
	var output []rune
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
		output = append(output, input)
	}
	return output, nil
}

func main() {
	info, _ := os.Stdin.Stat()

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: echo 'Hello' | rgbtext")
		os.Exit(2)
	}

	output, readErr := readCliInput()
	if readErr != nil {
		os.Exit(1)
	}
	rgb.Print(output)
}

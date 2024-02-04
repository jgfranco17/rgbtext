package main

import (
	"fmt"
	"os"

	reader "github.com/jgfranco17/rgbtext/core/pkg/reader"
	rgb "github.com/jgfranco17/rgbtext/core/pkg/rgb"
)

func main() {
	info, _ := os.Stdin.Stat()

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: echo 'Hello' | rgbtext")
		os.Exit(2)
	}

	output, readErr := reader.ReadCliInput()
	if readErr != nil {
		os.Exit(1)
	}
	rgb.Print(output)
}

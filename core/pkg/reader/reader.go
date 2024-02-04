package reader

import (
	"bufio"
	"io"
	"os"
)

func ReadCliInput() ([]rune, error) {
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

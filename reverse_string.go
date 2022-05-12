package reverse_string

import (
	"strings"
)

func reverse(input []rune) (result []rune) {
	result = make([]rune, len(input))
	copy(result, input)

	iLimit := len(input) / 2
	for i := 0; i < iLimit; i++ {
		idx := len(input) - 1 - i
		result[i], result[idx] = result[idx], result[i]
	}

	return
}

func reverseLine(line string) (lineOut string) {
	runes := []rune(line)

	lineOut = string(reverse(runes))

	return
}

func ReverseString(input string) (output string) {
	lines := strings.Split(input, "\n")
	for lineIdx := range lines {
		line := lines[lineIdx]

		lineReversed := reverseLine(line)

		lines[lineIdx] = lineReversed
	}

	output = strings.Join(lines, "\n")

	// solution goes here
	return
}

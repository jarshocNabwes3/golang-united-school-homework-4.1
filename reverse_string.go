package reverse_string

func reverse[T any](input []T) (result []T) {
	result = make([]T, len(input))
	copy(result, input)

	iLimit := len(input) / 2
	for i := 0; i < iLimit; i++ {
		idx := len(input) - 1 - i
		result[i], result[idx] = result[idx], result[i]
	}

	return
}

func ReverseString(input string) (output string) {
	runes := []rune(input)

	output = string(reverse(runes))

	// solution goes here
	return
}

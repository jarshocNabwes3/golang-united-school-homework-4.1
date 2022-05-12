package reverse_string

import (
	"testing"

	"gotest.tools/assert"
)

func TestReverseString(t *testing.T) {
	input := "Hello World\n" + " Worлd Heллo"
	rvOutputExpected := "dlroW olleH\n" + "oллeH dлroW "

	rvOutput := ReverseString(input)

	assert.Equal(t, rvOutputExpected, rvOutput, "Unexpected result\tExpected: %v\n\tGot: %v",
		rvOutputExpected, rvOutput)
}

package slices_test

import (
	"testing"

	"github.com/myazahq/go-test-exercise/section1/slices"
	"github.com/stretchr/testify/assert"
)

func TestSumOfEvenNumbers(t *testing.T){
	testCases := []struct{
		name string
		input []int
		output int
	}{
		{
			name: "empty slice",
			input: []int{},
			output: 0,
		},
		{
			name: "odd and even numbers",
			input: []int{1,2,3,4,4,7},
			output: 10,
		},
		{
			name: "even numbers only",
			input: []int{2, 6, 18},
			output: 26,
		},
		{
			name: "odd numbers only",
			input: []int{1,3,7},
			output: 0,
		},
	}

	for _, tc := range testCases{
		t.Run(tc.name, func(t *testing.T) {
			sum := slices.SumOfEvenNumbers(tc.input)
			assert.Equal(t, tc.output, sum)
		})
	}
}
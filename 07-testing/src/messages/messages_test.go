package messages

import "testing"

func Test_Sum(t *testing.T) {
	testCases := []struct {
		name   string
		input  Value
		output int64
	}{{
		name:   "Testing 1",
		input:  Value(1),
		output: 11,
	},
		{
			name:   "Testing 1000",
			input:  Value(1000),
			output: 1010,
		},
		{
			name:   "Testing 3",
			input:  Value(3),
			output: 13,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// we use the elipsis here to pass individual array elements
			result := tc.input.SumTen()
			if result != tc.output {
				t.Errorf("expected %d but got %d", tc.output, result)
			}
		})
	}
}

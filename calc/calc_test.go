package main

import (
	"testing"
)

func TestCalc(t *testing.T) {
	var cases = []struct {
		expected string
		input    string
	}{
		{
			input:    "1 2 3 4 + * + =",
			expected: "15",
		},
		{
			input:    "1 2 + 3 4 + * =",
			expected: "21",
		},
		{
			input:    "   1  2 \n   + \n   =  \n ",
			expected: "3",
		},
		{
			input:    "",
			expected: "",
		},
		{
			input:    "1 a + =",
			expected: "calc: strconv.Atoi: parsing \"a\": invalid syntax",
		},
		{
			input:    "1 2 + + =",
			expected: "calc: stack is empty",
		},
	}

	for _, item := range cases {
		if err, ok := Test(item.input, item.expected); !ok {
			if err != item.expected {
				t.Error("test for OK Failed - " + err)
			}
		}
	}
}

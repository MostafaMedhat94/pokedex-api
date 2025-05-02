package main

import (
	"testing"
)

type testCase struct {
	input    string
	expected []string
}

func TestCleanInput(t *testing.T) {

	cases := []testCase{
		{
			input:    "Sabahu Al Khair",
			expected: []string{"sabahu", "al", "khair"},
		},
		{
			input:    "Sabahu    Al Khair",
			expected: []string{"sabahu", "al", "khair"},
		},
		{
			input:    "   Massa Al Khair  ",
			expected: []string{"massa", "al", "khair"},
		},
		{
			input:    "  MASSA Al Khair  ",
			expected: []string{"massa", "al", "khair"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Expected %v, got %v", c.expected, actual)

			t.Fail()
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Expected %v, got %v", c.expected, actual)

				t.Fail()
			}
		}
	}

}

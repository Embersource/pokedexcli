package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "sadf awdf",
			expected: []string{"sadf", "awdf"},
		},
	}

	for _, v := range cases {
		actual := cleanInput(v.input)
		if len(actual) != len(v.expected) {
			t.Errorf("empty string")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := v.expected[i]
			if word != expectedWord {
				t.Errorf("%v", word)
			}
		}
	}
}

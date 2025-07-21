package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "pikachu charmander",
			expected: []string{"pikachu", "charmander"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("length of actual (%v) does not match expected length (%v)", len(actual), len(c.expected))
		}
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("actual word %v:%v does not match expected word %v:%v", i, actual[i], i, c.expected[i])
			}
		}
	}
}

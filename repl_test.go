package main

import (
	//  "fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "I liKe Redbull!",
			expected: []string{"i", "like", "redbull!"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHu",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("length of %v slice doesn't match expected: %v", actual, c.expected)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("word %v in slice %v does not match %v", word, actual, expectedWord)
			}
		}
	}
}

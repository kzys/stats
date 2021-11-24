package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBar(t *testing.T) {
	testcases := []struct {
		input    int
		expected string
	}{
		{0, " "},
		{1, "▏"},
		{10, "█▎"},
	}
	for _, tc := range testcases {
		assert.Equal(t, tc.expected, bar(tc.input))
		assert.Equal(t, tc.expected, bar(tc.input))
	}
}
func TestHistogram(t *testing.T) {
	testcases := []struct {
		input    []float64
		expected []int
	}{
		{
			[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
	}
	for _, tc := range testcases {
		bins, _, _ := histogram(tc.input, tc.input[0], tc.input[len(tc.input)-1])
		assert.Equal(t, tc.expected, bins)
	}
}

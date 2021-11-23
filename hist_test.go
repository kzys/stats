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

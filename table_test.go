package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZero(t *testing.T) {
	tw := &tableWriter{}
	lines, w := tw.Flush()
	assert.Nil(t, lines)
	assert.Equal(t, 0, w)
}

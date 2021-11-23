package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintSummary(t *testing.T) {
	buf := bytes.NewBufferString("")
	printSummary(buf, []float64{0.0, 1.0, 2.0})
	assert.Equal(t, `avg: 1.00
min: 0.00
p50: 1.00
p90: 1.00
p99: 1.00
max: 2.00
`, buf.String())
}

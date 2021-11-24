package main

import (
	"fmt"
	"strings"
)

func bar(x int) string {
	s := strings.Repeat("\u2588", x/8)
	mod := x % 8
	if mod == 0 {
		return s + " "
	}
	c := 0x258f - (mod - 1)
	return s + fmt.Sprintf("%c", c)
}

func printHistogram(values []float64, min float64, max float64) {
	binCount := 10
	bins := make([]int, binCount)
	binSize := (max - min + 1) / float64(binCount)
	end := min + binSize
	var index int
	for _, v := range values {
		if v >= end {
			index++
			end += binSize
		}
		bins[index]++
	}

	fmt.Printf("\nhistogram (binSize=%.2f, bins=%d)\n", binSize, binCount)
	for i, v := range bins {
		begin := min + float64(i)*binSize
		end := min + float64(i+1)*binSize
		fmt.Printf("[%.2f, %.2f) %3d: %s\n", begin, end, v, bar(v))
	}
}

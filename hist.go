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
	binSize := (max - min) / float64(binCount)
	end := binSize
	var index int
	var maxBin int
	for _, v := range values {
		if v > end {
			end += binSize
			index++
			continue
		}
		bins[index]++
		if bins[index] > maxBin {
			maxBin = bins[index]
		}
	}

	fmt.Printf("\nhistogram (bins=%d)\n", binCount)
	for _, v := range bins {
		fmt.Printf("%3d: %s\n", v, bar(v))
	}
}

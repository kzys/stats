package main

import (
	"fmt"
	"strings"
)

const binCount = 10

func bar(x int) string {
	s := strings.Repeat("\u2588", x/8)
	mod := x % 8
	if mod == 0 {
		return s + " "
	}
	c := 0x258f - (mod - 1)
	return s + fmt.Sprintf("%c", c)
}

func histogram(values []float64, min float64, max float64) ([]int, float64, int) {
	bins := make([]int, binCount)
	binSize := (max - min + 1) / float64(binCount)
	end := min + binSize
	var index int
	var maxBin int
	for _, v := range values {
		for v >= end {
			index++
			end += binSize
		}
		bins[index]++
		if bins[index] > maxBin {
			maxBin = bins[index]
		}
	}
	return bins, binSize, maxBin
}

func printHistogram(values []float64, min float64, max float64) {
	bins, binSize, maxBin := histogram(values, min, max)
	fmt.Printf("\nhistogram (binSize=%.2f, bins=%d)\n", binSize, binCount)

	tw := &tableWriter{}
	tw.Style(left, right, left, right, left, right, left)
	for i, v := range bins {
		begin := min + float64(i)*binSize
		end := min + float64(i+1)*binSize
		tw.Write(
			"[",
			fmt.Sprintf("%.2f", begin),
			", ",
			fmt.Sprintf("%.2f", end),
			") ",
			fmt.Sprintf("%d ", v),
		)
	}
	lines, width := tw.Flush()

	scale := float64(maxBin/8) / float64(80-width)
	if scale < 1.0 {
		scale = 1.0
	}

	for index, line := range lines {
		w := float64(bins[index]) / scale

		// Non-zero bins should be different from zero bins,
		// even it looks size-wise inaccurate.
		if bins[index] > 0 && int(w) == 0 {
			w = 1
		}
		fmt.Printf("%s%s\n", line, bar(int(w)))
	}
}

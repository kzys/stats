package main

import "fmt"

func printHistogram(values []float64, min float64, max float64) {
	binCount := 10
	bins := make([]int, binCount)
	binSize := (max - min) / float64(binCount)
	end := binSize
	var index int
	for _, v := range values {
		if v > end {
			bins = append(bins, 0)
			end += binSize
			index++
			continue
		}
		bins[index]++
	}
	fmt.Printf("bins = %v, binSize = %f\n", bins, binSize)
}

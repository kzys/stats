package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"

	"github.com/spf13/pflag"
)

func main() {
	ws := regexp.MustCompile(`\s+`)

	key := pflag.IntP("key", "k", 0, "key")
	pflag.Parse()

	rows := make([][]string, 0)

	r := bufio.NewReader(os.Stdin)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		columns := ws.Split(line, -1)
		rows = append(rows, columns)
	}

	var sum float64
	values := make([]float64, 0)

	if *key != 0 {
		index := *key - 1
		for y := 0; y < len(rows); y++ {
			columns := rows[y]
			if index >= len(columns) {
				continue
			}
			f, err := strconv.ParseFloat(rows[y][index], 64)
			if err != nil {
				continue
			}
			values = append(values, f)
			sum += f
		}
	}

	sort.Float64Slice(values).Sort()
	min := values[0]
	max := values[len(values)-1]

	fmt.Printf("avg: %.2f\n", sum/float64(len(values)))
	fmt.Printf("min: %.2f\n", min)

	percentiles := []float64{0.5, 0.9, 0.99}
	for _, p := range percentiles {
		index := float64(len(values)-1) * p
		fmt.Printf("p%.f: %.2f\n", p*100, values[int(index)])
	}
	fmt.Printf("max: %.2f\n", max)

	printHistogram(values, min, max)
}

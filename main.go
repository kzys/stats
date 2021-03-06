package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"

	"github.com/spf13/pflag"
)

func printSummary(f io.Writer, values []float64) {
	var sum float64

	for _, f := range values {
		sum += f
	}

	min := values[0]
	max := values[len(values)-1]

	fmt.Fprintf(f, "avg: %.2f\n", sum/float64(len(values)))
	fmt.Fprintf(f, "min: %.2f\n", min)

	percentiles := []float64{0.5, 0.9, 0.99}
	for _, p := range percentiles {
		index := float64(len(values)-1) * p
		fmt.Fprintf(f, "p%.f: %.2f\n", p*100, values[int(index)])
	}
	fmt.Fprintf(f, "max: %.2f\n", max)
}

func realMain() error {
	ws := regexp.MustCompile(`\s+`)

	key := pflag.IntP("key", "k", 0, "key")
	re := pflag.StringP("regexp", "r", "", "regexp")
	pflag.Parse()

	var pattern *regexp.Regexp
	if *re != "" {
		var err error
		pattern, err = regexp.Compile(*re)
		if err != nil {
			return err
		}
		*key = 1
	}

	rows := make([][]string, 0)

	r := bufio.NewReader(os.Stdin)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}

		if pattern != nil {
			matches := pattern.FindStringSubmatch(line)
			rows = append(rows, matches)
		} else {
			columns := ws.Split(line, -1)
			rows = append(rows, columns)
		}
	}

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
		}
	} else {
		pflag.Usage()
		os.Exit(1)
	}

	sort.Float64Slice(values).Sort()

	min := values[0]
	max := values[len(values)-1]

	printSummary(os.Stdout, values)
	printHistogram(values, min, max)

	return nil
}

func main() {
	err := realMain()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err.Error())
		os.Exit(1)
	}
}

package main

import (
	"strings"
)

const (
	left = iota
	right
)

type tableWriter struct {
	rows  [][]string
	style []int
}

func (tw *tableWriter) Style(style ...int) {
	tw.style = style
}

func (tw *tableWriter) Write(columns ...string) {
	tw.rows = append(tw.rows, columns)
}

func (tw *tableWriter) Flush() ([]string, int) {
	max := make(map[int]int)

	for _, columns := range tw.rows {
		for i, v := range columns {
			len := len(v)
			if len > max[i] {
				max[i] = len
			}
		}
	}

	var lines []string
	width := 0
	for _, columns := range tw.rows {
		var line string
		for i, v := range columns {
			len := len(v)
			if tw.style[i] == right {
				line += strings.Repeat(" ", max[i]-len)
			}
			line += v
		}
		if len(line) > width {
			width = len(line)
		}
		lines = append(lines, line)
	}
	return lines, width
}

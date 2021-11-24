package main

import (
	"fmt"
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

func (tw *tableWriter) Flush() {
	max := make(map[int]int)

	for _, columns := range tw.rows {
		for i, v := range columns {
			len := len(v)
			if len > max[i] {
				max[i] = len
			}
		}
	}

	for _, columns := range tw.rows {
		for i, v := range columns {
			len := len(v)
			if tw.style[i] == right {
				fmt.Print(strings.Repeat(" ", max[i]-len))
			}
			fmt.Print(v)
		}
		fmt.Println()
	}
}

package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

type vent struct {
	x1, y1 int
	x2, y2 int
}

var chart [1000][1000]int

func drawLine(v vent) {
	switch {
	case v.x1 == v.x2:
		drawVertical(v)
	case v.y1 == v.y2:
		drawHorizontal(v)
	default:
		drawDiagonal(v)
	}
}

func drawDiagonal(v vent) {

	x1, y1, x2, y2 := v.x1, v.y1, v.x2, v.y2
	if x1 > x2 {
		// make sure x goes small -> large
		x1, y1, x2, y2 = x2, y2, x1, y1
	}

	// is y rising or falling?
	riseFall := 1
	if y1 > y2 {
		riseFall = -1
	}

	i, j := x1, y1
	for i <= x2 {
		chart[i][j]++

		i++
		j += riseFall
	}

}

func drawVertical(v vent) {

	y1, y2 := v.y1, v.y2
	if y1 > y2 {
		y1, y2 = y2, y1
	}

	for i := y1; i <= y2; i++ {
		chart[v.x1][i]++
	}
}

func drawHorizontal(v vent) {

	x1, x2 := v.x1, v.x2
	if x1 > x2 {
		x1, x2 = x2, x1
	}

	for i := x1; i <= x2; i++ {
		chart[i][v.y1]++
	}
}

func countPointsAtLeast(v int) int {
	c := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if chart[i][j] >= v {
				c++
			}
		}
	}

	return c
}

func run(args []string, stdout io.Writer) error {

	for _, v := range vents {
		drawLine(v)
	}

	log.Printf("there are %d very dangerous points", countPointsAtLeast(2))

	return nil
}

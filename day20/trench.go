package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		log.Fatalf("program failed: %s\n", err)
	}

	// fmt.Printf("fini\n")
}

func run(args []string, stdout io.Writer) error {

	// log.Printf("algo is %d long", len(algo))
	// log.Printf("testAlgo is %d long", len(testAlgo))

	// i := img(testImage)
	// for k := 0; k < 50; k++ {
	// 	i = i.pad(".", 4)
	// 	i = i.apply(testAlgo)
	// 	i = i.trim()
	// }

	i := img(image)
	for k := 0; k < 25; k++ {
		i = i.pad(".", 4)
		i = i.apply(algo)
		i = i.trim().pad("#", 4)
		i = i.apply(algo)
		i = i.trim()
	}

	i.print()
	log.Printf("lit pixel count %d", i.countLit())

	return nil
}

type img []string

func (i img) pad(c string, n int) img {
	p := strings.Repeat(c, len(i[0])+(2*n))

	ni := []string{}

	for i := 0; i < n; i++ {
		ni = append(ni, p)
	}

	for _, l := range i {
		ni = append(ni, strings.Repeat(c, n)+l+strings.Repeat(c, n))
	}

	for i := 0; i < n; i++ {
		ni = append(ni, p)
	}

	return ni
}

func (i img) trim() img {

	firstY, lastY := -1, math.MaxInt
	minLeft, maxRight := math.MaxInt, 0

	for y := 0; y < len(i); y++ {

		ind := strings.Index(i[y], "#")
		if ind == -1 {
			continue
		}

		rind := strings.LastIndex(i[y], "#")

		if firstY == -1 {
			firstY = y
		}
		lastY = y

		if ind < minLeft {
			minLeft = ind
		}
		if rind > maxRight {
			maxRight = rind
		}
	}

	ni := img{}

	for y := firstY; y <= lastY; y++ {
		ni = append(ni, i[y][minLeft:maxRight+1])
	}

	return ni
}

func (i img) apply(algo string) img {

	ni := []string{}

	for y := 1; y < len(i)-1; y++ {
		nl := ""
		for x := 1; x < len(i[0])-1; x++ {
			o := i.getOffset(x, y)
			c := algo[o]
			nl += string(c)
		}
		ni = append(ni, nl)
	}

	return ni
}

func (i img) getOffset(x, y int) int {

	n := i[y-1][x-1:x+2] + i[y][x-1:x+2] + i[y+1][x-1:x+2]
	n = strings.Map(func(c rune) rune {
		if c == '#' {
			return '1'
		}
		return '0'
	}, n)

	nn, err := strconv.ParseInt(n, 2, 0)
	if err != nil {
		log.Fatalf("failed to parse offset %s: %v", n, err)
	}

	return int(nn)
}

func (i img) countLit() int {
	c := 0
	for y := 0; y < len(i); y++ {
		for x := 0; x < len(i[0]); x++ {
			if i[y][x] == '#' {
				c++
			}
		}
	}
	return c
}

func (i img) print() {
	for y := 0; y < len(i); y++ {
		fmt.Printf("%s\n", i[y])
	}
	fmt.Printf("%s\n", strings.Repeat("-", len(i[0])))
}

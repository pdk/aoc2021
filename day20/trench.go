package main

import (
	"fmt"
	"io"
	"log"
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

	// i := img(testImage).pad().pad()
	// i.print()
	// i = i.apply(testAlgo)
	// i.print()
	// i = i.apply(testAlgo)
	// i.print()

	i := img(image).pad(".").pad(".")
	i.print()
	// log.Printf("start with %d pixels", i.countLit())
	i = i.apply(algo)
	i.print()
	i = i.pad("#").pad("#")
	i = i.apply(algo)
	i.print()

	// log.Printf("first apply gets %d pixels", i.countLit())
	// i = i.apply(algo)
	log.Printf("second apply gets %d pixels", i.countLit())

	return nil
}

type img []string

func (i img) pad(c string) img {
	p := strings.Repeat(c, len(i[0])+4)

	ni := []string{p, p}

	for _, l := range i {
		ni = append(ni, c+c+l+c+c)
	}

	ni = append(ni, p, p)

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

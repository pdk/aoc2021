package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		log.Fatalf("program failed: %s\n", err)
	}
	fmt.Printf("fini\n")
}

func run(args []string, stdout io.Writer) error {

	onCount := 0

	for i := -50; i <= 50; i++ {
		for j := -50; j <= 50; j++ {
			for k := -50; k <= 50; k++ {
				if isOn(i, j, k) {
					onCount++
				}
			}
		}
	}

	log.Printf("%d cells are on", onCount)

	return nil
}

func isOn(x, y, z int) bool {
	for i := len(data) - 1; i >= 0; i-- {
		e := data[i].effect(x, y, z)
		if e != nada {
			return e == on
		}
	}
	return false
}

type which int

const (
	nada which = iota
	on
	off
)

type lohi struct {
	from, to int
}

type cmd struct {
	which
	xRange, yRange, zRange lohi
}

func (c cmd) effect(x, y, z int) which {
	if x >= c.xRange.from && x <= c.xRange.to &&
		y >= c.yRange.from && y <= c.yRange.to &&
		z >= c.zRange.from && z <= c.zRange.to {

		return c.which
	}

	return nada
}

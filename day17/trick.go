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

	// t := 0
	// for i := 0; i < 100; i++ {
	// 	log.Printf("i=%d t=%d", i, t)
	// 	t += i
	// 	if t > 171 {
	// 		break
	// 	}
	// }

	// target area: x=137..171, y=-98..-73
	target := newArea(137, 171, -98, -73)
	success := false

	for tryY := 1000; tryY > 10; tryY-- {

		p := pos{0, 0}
		v := velocity{17, tryY}
		highest := 0

		for i := 0; i < 200000; i++ {

			p, v = step(p, v)
			highest = max(highest, p.y)

			if i%1000 == 0 {
				log.Printf("i=%6d p=%3d,%3d v=%3d,%3d in=%t below=%t", i, p.x, p.y, v.x, v.y, isWithin(p, target), isBelow(p, target))
			}

			if isBelow(p, target) {
				log.Printf("missed it")
				break
			}

			if isWithin(p, target) {
				success = true
				log.Printf("hit!!!")
				log.Printf("i=%6d p=%3d,%3d v=%3d,%3d in=%t below=%t", i, p.x, p.y, v.x, v.y, isWithin(p, target), isBelow(p, target))
				break
			}
		}
		if success {
			log.Printf("tryY = %d, highest = %d", tryY, highest)
			break
		}
	}

	return nil
}

type area struct {
	nw, se pos
}

func newArea(x1, x2, y1, y2 int) area {
	return area{
		nw: pos{
			min(x1, x2),
			max(y1, y2),
		},
		se: pos{
			max(x1, x2),
			min(y1, y2),
		},
	}
}

type pos struct {
	x, y int
}

type velocity struct {
	x, y int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func isWithin(p pos, a area) bool {

	return p.x >= a.nw.x &&
		p.y <= a.nw.y &&
		p.x <= a.se.x &&
		p.y >= a.se.y

}

func isBelow(p pos, a area) bool {
	return p.y < a.se.y
}

// step computes a new position and velocity
func step(p pos, v velocity) (pos, velocity) {

	p.x += v.x
	p.y += v.y

	v.x += drag(v.x)
	v.y-- // gravity

	return p, v
}

func drag(x int) int {
	if x > 0 {
		return -1
	}
	if x < 0 {
		return 1
	}
	return 0
}

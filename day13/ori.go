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

	// log.Printf("example has %d dots, size %#v", len(example), example.maxPt())

	// example.print()
	// fmt.Println("-------------------------")

	// // fold along y=7
	// // fold along x=5

	// f := example.foldAtY(7)
	// f.print()
	// fmt.Println("-------------------------")
	// f = f.foldAtX(5)
	// f.print()

	log.Printf("start has %d dots, size %#v", len(start), start.maxPt())

	f := start.foldAtX(655)
	f = f.foldAtY(447)
	f = f.foldAtX(327)
	f = f.foldAtY(223)
	f = f.foldAtX(163)
	f = f.foldAtY(111)
	f = f.foldAtX(81)
	f = f.foldAtY(55)
	f = f.foldAtX(40)
	f = f.foldAtY(27)
	f = f.foldAtY(13)
	f = f.foldAtY(6)

	log.Printf("folded has %d dots, size %#v", len(f), f.maxPt())

	f.print()

	return nil
}

type sheet map[pt]bool

type pt struct {
	x, y int
}

func (s sheet) print() {
	size := s.maxPt()

	for j := 0; j <= size.y; j++ {
		for i := 0; i <= size.x; i++ {
			if s[pt{i, j}] {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func (s sheet) maxPt() pt {
	m := pt{}

	for k := range s {
		if k.x > m.x {
			m.x = k.x
		}
		if k.y > m.y {
			m.y = k.y
		}
	}

	return m
}

func (s sheet) foldAtX(at int) sheet {
	folded := sheet{}

	for k := range s {
		if k.x > at {
			k.x = 2*at - k.x
		}
		folded[k] = true
	}

	return folded
}

func (s sheet) foldAtY(at int) sheet {
	folded := sheet{}

	for k := range s {
		if k.y > at {
			k.y = 2*at - k.y
		}
		folded[k] = true
	}

	return folded
}

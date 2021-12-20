package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		log.Fatalf("program failed: %s\n", err)
	}
	fmt.Printf("fini\n")
}

func run(args []string, stdout io.Writer) error {

	loadCaves()

	for c := range caves {
		if isLittleCave(c) && c != "start" {
			spelunk([]string{"start"}, c)
		}
	}

	log.Printf("there are %d paths thru these caves", len(allPaths))

	return nil
}

var allPaths = map[string]bool{}

func savePath(path []string) {
	p := strings.Join(path, "->")
	if !allPaths[p] {
		log.Printf("found a new path: %s", p)
		allPaths[p] = true
	}
}

func spelunk(path []string, oneSmallCave string) {

	if len(path) > 1000 {
		log.Printf("aborting spelunk, path too long: %#v", path)
		return
	}

	here := path[len(path)-1]

	if here == "end" {
		savePath(path)
		return
	}

	if isLittleCave(here) {
		visits := visitCount(here, path)

		if here == oneSmallCave {
			if visits > 2 {
				return
			}
		} else {
			if visits > 1 {
				return
			}
		}
	}

	for connected := range caves[here].connected {
		spelunk(append(path, connected), oneSmallCave)
	}
}

func isLittleCave(name string) bool {
	return name == strings.ToLower(name)
}

func visitCount(name string, path []string) int {
	c := 0
	for _, each := range path {
		if name == each {
			c++
		}
	}
	return c
}

func loadCaves() {
	for _, conn := range data {
		connect(conn.here, conn.there)
		connect(conn.there, conn.here)
	}
}

func connect(c1, c2 string) {
	cv, defined := caves[c1]
	if !defined {
		cv = &cave{
			name:      c1,
			connected: map[string]bool{},
		}
		caves[c1] = cv
	}
	cv.connected[c2] = true
}

// cave is a named space with a set of connected caves (names)
type cave struct {
	name      string
	connected map[string]bool
}

// caves maps from a cave to other connected caves
var caves = map[string]*cave{}

// connection is just to frame data loading
type connection struct {
	here, there string
}

var data = []connection{
	{"EG", "bj"},
	{"LN", "end"},
	{"bj", "LN"},
	{"yv", "start"},
	{"iw", "ch"},
	{"ch", "LN"},
	{"EG", "bn"},
	{"OF", "iw"},
	{"LN", "yv"},
	{"iw", "TQ"},
	{"iw", "start"},
	{"TQ", "ch"},
	{"EG", "end"},
	{"bj", "OF"},
	{"OF", "end"},
	{"TQ", "start"},
	{"TQ", "bj"},
	{"iw", "LN"},
	{"EG", "ch"},
	{"yv", "iw"},
	{"KW", "bj"},
	{"OF", "ch"},
	{"bj", "ch"},
	{"yv", "TQ"},
}

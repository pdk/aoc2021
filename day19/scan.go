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

	merged, ok := mergeAll(data[0], data[1:])

	log.Printf("merged is %t, size %d", ok, merged.size())

	return nil
}

func mergeAll(base scanner, toMerge []scanner) (scanner, bool) {

	c := 0
	for len(toMerge) > 0 && c < 100 {
		c++

		next := toMerge[0]

		matching, f, r, p := base.findNMatches(12, next)
		if len(matching) != 0 {
			next = next.rotate(f, r).shift(p)
			base = base.merge(next)
			toMerge = toMerge[1:]
		} else {
			// put this one back on the end
			toMerge = append(toMerge[1:], next)
		}
	}

	return base, len(toMerge) == 0
}

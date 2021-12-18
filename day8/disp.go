package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

// 1 => 2 segs
// 7 => 3 segs
// 4 => 4 segs
// 8 => 7 segs

// 2 => 5 segs
// 3 => 5 segs
// 5 => 5 segs
// 0 => 6 segs
// 6 => 6 segs
// 9 => 6 segs

func run(args []string, stdout io.Writer) error {

	decodeCount := 0

	for _, d := range displayNotes {
		decoder := identify(d.distinct)

		decodeCount += countMatch(d.display, decoder)

		// log.Printf("sample %d: %#v", i, decoder)

		// if i > 10 {
		// 	break
		// }
	}

	log.Printf("decode count %d", decodeCount)

	return nil
}

func countMatch(disp [4]string, decoder map[string]int) int {

	c := 0

	for _, d := range disp {
		d = canon(d)
		_, ok := decoder[d]
		if ok {
			c++
		}
	}

	return c
}

func identify(sample [10]string) map[string]int {

	decoder := map[string]int{}

	for _, s := range sample {
		s = canon(s)
		switch len(s) {
		case 2:
			decoder[s] = 1
		case 3:
			decoder[s] = 7
		case 4:
			decoder[s] = 4
		case 7:
			decoder[s] = 8
		}
	}

	return decoder
}

func canon(s string) string {
	letters := strings.Split(s, "")
	sort.Strings(letters)
	return strings.Join(letters, "")
}

package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if err := part1(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	if err := part2(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func part1(args []string, stdout io.Writer) error {

	count := 0
	last := data[0]
	for _, datum := range data[1:] {
		if datum > last {
			count++
		}
		last = datum
	}

	log.Printf("part 1 count = %d", count)

	return nil
}

func part2(args []string, stdout io.Writer) error {

	count := 0

	for i := 0; i < len(data)-3; i++ {
		last := data[i] + data[i+1] + data[i+2]
		this := data[i+1] + data[i+2] + data[i+3]

		if this > last {
			count++
		}
	}

	log.Printf("part 2 count = %d", count)

	return nil
}

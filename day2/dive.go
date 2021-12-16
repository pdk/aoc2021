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

	hori := 0
	depth := 0

	for _, next := range data {
		switch next.dim {
		case "forward":
			hori += next.amount
		case "up":
			depth -= next.amount
		case "down":
			depth += next.amount
		default:
			return fmt.Errorf("don't know navigation command %s", next.dim)
		}

	}

	log.Printf("part 1 horiontal %d, depth %d, navpoint %d", hori, depth, hori*depth)

	return nil
}

func part2(args []string, stdout io.Writer) error {

	aim := 0
	hori := 0
	depth := 0

	for _, next := range data {
		switch next.dim {
		case "forward":
			hori += next.amount
			depth += aim * next.amount
		case "up":
			aim -= next.amount
		case "down":
			aim += next.amount
		default:
			return fmt.Errorf("don't know navigation command %s", next.dim)
		}

	}

	log.Printf("part 2 horiontal %d, depth %d, navpoint %d", hori, depth, hori*depth)

	return nil
}

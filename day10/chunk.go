package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	if err := part1(os.Args, os.Stdout); err != nil {
		log.Fatalf("program failed: %s\n", err)
	}
	log.Printf("fini\n")
}

func part1(args []string, stdout io.Writer) error {

	score := 0

	for n, line := range data {

		nextScore := parseChunks(line)
		if nextScore > 0 {
			log.Printf("%d score! %d", n, nextScore)
			score += nextScore
		}

	}

	log.Printf("total score %d", score)

	return nil
}

var (
	errLineIncomplete = fmt.Errorf("line is incomplete")
	errLineCorrupted  = fmt.Errorf("line is corrupted")

	matcher = map[rune]rune{
		'(': ')',
		'[': ']',
		'<': '>',
		'{': '}',
	}
)

func parseChunks(input string) int {

	for len(input) > 0 {
		var err error
		input, err = parse(input)
		if err == errLineIncomplete {
			return 0
		}
		if err == errLineCorrupted {
			switch input {
			case ")":
				return 3
			case "}":
				return 1197
			case "]":
				return 57
			case ">":
				return 25137
			}
		}
	}

	return 0
}

func parse(input string) (string, error) {
	var err error

	if len(input) == 0 {
		return input, nil
	}

	next := rune(input[0])
	remaining := input[1:]

	match, ok := matcher[next]
	if !ok {
		return string(next), errLineCorrupted
	}

	for {
		if len(remaining) == 0 {
			return remaining, errLineIncomplete
		}

		if strings.HasPrefix(remaining, string(match)) {
			return remaining[1:], nil
		}

		remaining, err = parse(remaining)
		if err != nil {
			return remaining, err
		}
	}
}

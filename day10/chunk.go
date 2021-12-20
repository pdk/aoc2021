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
	// if err := part1(os.Args, os.Stdout); err != nil {
	// 	log.Fatalf("program failed: %s\n", err)
	// }
	if err := part2(os.Args, os.Stdout); err != nil {
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

func part2(args []string, stdout io.Writer) error {

	scores := []int{}

	for n, line := range data {
		nextScore := completeChunks(line)
		if nextScore > 0 {
			log.Printf("%d score! %d", n, nextScore)
			scores = append(scores, nextScore)
		}
	}

	log.Printf("got %d scores", len(scores))

	sort.Ints(scores)
	p := len(scores) / 2

	log.Printf("middle score %d is %d", p, scores[p])

	return nil
}

func score(s string) int {
	// 	Did you know that autocomplete tools also have contests? It's true! The
	// 	score is determined by considering the completion string
	// 	character-by-character. Start with a total score of 0. Then, for each
	// 	character, multiply the total score by 5 and then increase the total
	// 	score by the point value given for the character in the following table:

	// ): 1 point.
	// ]: 2 points.
	// }: 3 points.
	// >: 4 points.

	score := 0

	for _, c := range s {

		score *= 5

		switch c {
		case ')':
			score += 1
		case ']':
			score += 2
		case '}':
			score += 3
		case '>':
			score += 4
		}
	}

	return score
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

func completeChunks(input string) int {

	working := input

	for len(working) > 0 {
		var err error
		working, err = parse(working)
		if err == errLineIncomplete {
			return score(working)
		}
		if err == errLineCorrupted {
			return 0
		}
	}

	return 0
}

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
			return string(match), errLineIncomplete
		}

		if strings.HasPrefix(remaining, string(match)) {
			return remaining[1:], nil
		}

		remaining, err = parse(remaining)
		if err != nil {
			return remaining + string(match), err
		}
	}
}

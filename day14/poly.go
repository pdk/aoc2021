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

	t := start

	for i := 0; i < 10; i++ {
		t = insert(t)
		log.Printf("counts %s", printable(counts(t)))
	}

	most, least := 0, 1000000
	for _, v := range counts(t) {
		if v > most {
			most = v
		}
		if v < least {
			least = v
		}
	}

	log.Printf("most %d, least %d, diff %d", most, least, most-least)

	return nil
}

func printable(counts map[rune]int) string {
	s := ""
	for k, v := range counts {
		s += fmt.Sprintf("%c: %d, ", k, v)
	}
	return s
}

func counts(template string) map[rune]int {
	r := map[rune]int{}

	for _, each := range template {
		r[each]++
	}

	return r
}

func insert(template string) string {
	r := ""

	var last rune
	for _, each := range template {

		ins, ok := rules[insertKey{last, each}]
		if ok {
			r += string(ins)
		}
		r += string(each)
		last = each
	}

	return r
}

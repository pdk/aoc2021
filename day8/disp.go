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

func run(args []string, stdout io.Writer) error {

	decodeCount := 0
	decodedSum := 0

	for i, d := range displayNotes {
		decoder := identify(d.distinct)

		decodeCount += countMatch(d.display, decoder)

		// log.Printf("sample %d: %#v", i, decoder)

		decodedValue := decode(decoder, d.display)

		log.Printf("decoded display: %d", decodedValue)

		decodedSum += decodedValue

		if i > 10000 {
			break
		}
	}

	log.Printf("decode count %d", decodeCount)
	log.Printf("sum decoded values: %d", decodedSum)

	return nil
}

func decode(decoder map[string]int, value [4]string) int {

	return 1000*decoder[canon(value[0])] +
		100*decoder[canon(value[1])] +
		10*decoder[canon(value[2])] +
		decoder[canon(value[3])]
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

//              a b c d e f g

// 1 => 2 segs:     c     f
// 7 => 3 segs, a   c     f
// 4 => 4 segs,   b c d   f
// 8 => 7 segs, a b c d e f g

// 2 => 5 segs, a   c d e   g
// 3 => 5 segs, a   c d   f g
// 5 => 5 segs, a b   d   f g

// 0 => 6 segs, a b c   e f g
// 6 => 6 segs, a b   d e f g
// 9 => 6 segs, a b c d   f g

// a 8 times
// b 6 times **
// c 8 times
// d 7 times
// e 4 times **
// f 9 times **
// g 7 times

func identify(sample [10]string) map[string]int {

	// log.Printf("identify: %s", sample)

	decoder := map[string]int{}
	recoder := [10]string{}

	for _, s := range sample {
		s = canon(s)
		switch len(s) {
		case 2:
			decoder[s] = 1
			recoder[1] = s
		case 3:
			decoder[s] = 7
			recoder[7] = s
		case 4:
			decoder[s] = 4
			recoder[4] = s
		case 7:
			decoder[s] = 8
			recoder[8] = s
		}
	}

	segA, segB, segC, segD, segE, segF, segG := ".", ".", ".", ".", ".", ".", "."

	// trans provides mapping from found segment to corrected segment.
	trans := map[string]string{}

	freq := frequency(sample)
	for seg, count := range freq {
		switch count {
		case 6:
			trans[seg] = "b"
			segB = seg
		case 4:
			trans[seg] = "e"
			segE = seg
		case 9:
			trans[seg] = "f"
			segF = seg
		}
	}

	candi := minus(recoder[7], recoder[1])
	if len(candi) != 1 {
		log.Fatalf("failed to find segA from 1 (%s), 7 (%s)", recoder[1], recoder[7])
	}

	segA = candi[0]
	trans[segA] = "a"

	known := segA + segB + segE + segF

	candi = minus(recoder[1], known)
	if len(candi) != 1 {
		log.Fatalf("failed to find segC")
	}

	segC = candi[0]
	trans[segC] = "c"

	known = segA + segB + segC + segE + segF

	candi = minus(recoder[4], known)
	if len(candi) != 1 {
		log.Fatalf("failed to find segD")
	}

	segD = candi[0]
	trans[segD] = "d"

	known = segA + segB + segC + segD + segE + segF

	candi = minus(recoder[8], known)
	if len(candi) != 1 {
		log.Fatalf("failed to find segG")
	}

	segG = candi[0]
	trans[segG] = "g"

	// log.Printf("trans %s %s %s %s %s %s %s", segA, segB, segC, segD, segE, segF, segG)

	decoder[canon(segA+segC+segD+segE+segG)] = 2
	decoder[canon(segA+segC+segD+segF+segG)] = 3
	decoder[canon(segA+segB+segD+segF+segG)] = 5
	decoder[canon(segA+segB+segC+segE+segF+segG)] = 0
	decoder[canon(segA+segB+segD+segE+segF+segG)] = 6
	decoder[canon(segA+segB+segC+segD+segF+segG)] = 9

	// log.Printf("decoder %v", decoder)

	return decoder
}

func canon(s string) string {
	letters := strings.Split(s, "")
	sort.Strings(letters)
	return strings.Join(letters, "")
}

func frequency(sample [10]string) map[string]int {

	freq := map[string]int{}

	for _, s := range sample {
		for _, c := range strings.Split(s, "") {
			freq[c]++
		}
	}

	// log.Printf("frequency %v", freq)

	return freq
}

// minus returns a slice of (single char) strings that appear in s1, but not in
// s2.
func minus(s1, s2 string) []string {

	missing := []string{}

	for _, x := range strings.Split(s1, "") {
		if strings.Index(s2, x) == -1 {
			missing = append(missing, x)
		}
	}

	return missing
}

// given two (canonical) strings, return what is NOT shared between them.
func extraOf(a, b string) string {
	diff := ""

	for _, l := range a {
		if strings.Index(b, string(l)) == -1 {
			diff += string(l)
		}
	}

	for _, l := range b {
		if strings.Index(a, string(l)) == -1 {
			diff += string(l)
		}
	}

	return diff
}

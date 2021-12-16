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

var masks = []int{
	//123456789012
	0b000000000001,
	0b000000000010,
	0b000000000100,
	0b000000001000,
	0b000000010000,
	0b000000100000,
	0b000001000000,
	0b000010000000,
	0b000100000000,
	0b001000000000,
	0b010000000000,
	0b100000000000,
}

func isBitSet(datum, pos int) bool {
	m := masks[pos]
	return (datum & m) != 0
}

// returns new value after setting bit
func setBitAt(value, pos int) int {
	return value | masks[pos]
}

func countBits(data []int, pos int) (int, int) {
	bit0 := 0
	bit1 := 0

	for _, d := range data {
		if isBitSet(d, pos) {
			bit1++
		} else {
			bit0++
		}
	}

	return bit0, bit1
}

func part1(args []string, stdout io.Writer) error {

	var gamma, epsilon int

	for p := 0; p < 12; p++ {

		bit0, bit1 := countBits(data, p)

		if bit1 > bit0 {
			gamma = setBitAt(gamma, p)
		} else {
			epsilon = setBitAt(epsilon, p)
		}
	}

	log.Printf("gamma %d (%b), epsilon %d (%b), power consumption %d", gamma, gamma, epsilon, epsilon, gamma*epsilon)

	return nil
}

func filterBitPosOn(data []int, pos int) []int {

	filtered := []int{}

	for _, d := range data {
		if isBitSet(d, pos) {
			filtered = append(filtered, d)
		}
	}

	return filtered
}

func filterBitPosOff(data []int, pos int) []int {

	filtered := []int{}

	for _, d := range data {
		if !isBitSet(d, pos) {
			filtered = append(filtered, d)
		}
	}

	return filtered
}

func part2(args []string, stdout io.Writer) error {

	o := oxygenGeneratorRating()
	c := co2ScrubberRating()

	log.Printf("oxygen generator rating %d, co2 scrubber rating %d, life support rating %d", o, c, o*c)

	return nil
}

func oxygenGeneratorRating() int {

	// make a local copy of the data
	data := append([]int{}, data...)

	for p := 11; p > 0; p-- {

		bit0, bit1 := countBits(data, p)

		if bit1 >= bit0 {
			data = filterBitPosOn(data, p)
		} else {
			data = filterBitPosOff(data, p)
		}

		if len(data) == 1 {
			break
		}
	}

	return data[0]
}

func co2ScrubberRating() int {

	// make a local copy of the data
	data := append([]int{}, data...)

	for p := 11; p > 0; p-- {

		bit0, bit1 := countBits(data, p)

		if bit1 < bit0 {
			data = filterBitPosOn(data, p)
		} else {
			data = filterBitPosOff(data, p)
		}

		if len(data) == 1 {
			break
		}
	}

	return data[0]
}

package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
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

func run(args []string, stdout io.Writer) error {

	var gamma, epsilon int

	for p := 0; p < 12; p++ {

		bit0 := 0
		bit1 := 0

		for _, d := range data {
			if isBitSet(d, p) {
				bit1++
			} else {
				bit0++
			}
		}

		log.Printf("pos %d: bit0 %d, bit1 %d", p, bit0, bit1)

		if bit1 > bit0 {
			gamma = setBitAt(gamma, p)
		} else {
			epsilon = setBitAt(epsilon, p)
		}
	}

	log.Printf("gamma %d (%b), epsilon %d (%b), power consumption %d", gamma, gamma, epsilon, epsilon, gamma*epsilon)

	return nil
}

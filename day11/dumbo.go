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

	loadGrid()
	doOneHundredSteps()

	log.Printf("flash count is %d", flashCount)

	return nil
}

type dumbo struct {
	energy  int
	flashed bool
}

var grid [10][10]dumbo

func loadGrid() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			grid[i][j].energy = data[i][j]
			grid[i][j].flashed = false // totally redundant
		}
	}
}

func doOneHundredSteps() {
	for i := 0; i < 100; i++ {
		doOneStep()
		displayGrid()
	}
}

func displayGrid() {
	// print out the 10x10 grid of energy levels
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("%2d", grid[i][j].energy)
		}
		fmt.Println()
	}
	fmt.Println("--------------------")
}

func doOneStep() {
	pumpItUp()
	flashThemAll()
	resetFlashersToZero()
}

func pumpItUp() {
	// increment all energy levels by 1
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			grid[i][j].energy++
		}
	}
}

func flashThemAll() {
	// any dumbo with energy > 9 flashes
	// incrementing neighbor energy levels by 1
	// repeat

	for {
		flashCount := 0
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if grid[i][j].energy > 9 && !grid[i][j].flashed {
					flash(i, j)
					flashCount++
				}
			}
		}
		if flashCount == 0 {
			return
		}
	}
}

var flashCount int

func flash(i, j int) {
	grid[i][j].flashed = true
	flashCount++

	incr(i-1, j-1)
	incr(i-1, j+0)
	incr(i-1, j+1)
	incr(i+0, j-1)
	incr(i+0, j+1)
	incr(i+1, j-1)
	incr(i+1, j+0)
	incr(i+1, j+1)
}

func incr(i, j int) {
	if i < 0 || i > 9 || j < 0 || j > 9 {
		return
	}
	grid[i][j].energy++
}

func resetFlashersToZero() {
	// set all that flashed to 0, reset flash flag
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if grid[i][j].flashed {
				grid[i][j].flashed = false
				grid[i][j].energy = 0
			}
		}
	}
}

var data = [10][10]int{
	{5, 6, 5, 1, 3, 4, 1, 4, 5, 2},
	{1, 3, 8, 1, 5, 4, 1, 2, 5, 2},
	{1, 8, 7, 8, 4, 3, 5, 2, 2, 4},
	{6, 8, 1, 4, 8, 3, 1, 5, 3, 5},
	{3, 8, 8, 3, 5, 4, 7, 3, 8, 3},
	{6, 4, 7, 3, 5, 4, 8, 4, 6, 4},
	{1, 8, 8, 5, 8, 3, 3, 6, 5, 8},
	{3, 7, 3, 2, 5, 8, 4, 7, 5, 2},
	{1, 8, 8, 1, 5, 4, 6, 1, 2, 8},
	{5, 1, 2, 1, 7, 1, 7, 7, 7, 6},
}

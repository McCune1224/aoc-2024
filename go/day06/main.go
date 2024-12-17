package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	Part1()
}

func Part1() {
	grid, err := ReadInput("./test_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// for _, v := range grid.rawMap {
	// 	fmt.Println(v)
	// }
	// fmt.Println(grid.guard.X, grid.guard.Y)

	var stopCheck error
	stopCheck = nil
	fmt.Println(stopCheck)
	for stopCheck == nil {
		stopCheck = grid.Step()
		print2DSlice(grid.trackedMap)
		time.Sleep(time.Second) // Optional delay to see changes
	}
}

func print2DSlice(slice [][]string) {
	// Move cursor up by number of rows
	fmt.Printf("\033[%dA\r", len(slice))
	for i := range slice {
		for j := range slice[i] {
			fmt.Printf("%s ", slice[i][j])
		}
		fmt.Print("\n")
	}
}

type Guard struct {
	Glyph     string
	Direction Direction
	X         int
	Y         int
}

type Direction struct {
	Xstep int
	Ystep int
}

var (
	stepLeft       = Direction{-1, 0}
	stepRight      = Direction{1, 0}
	stepUp         = Direction{0, 1}
	stepDown       = Direction{0, -1}
	ErrOutOfBounds = errors.New("OUT OF BOUNDS")
)

func (g *Grid) Turn() {
	switch g.guard.Glyph {
	case "^":
		g.guard.Glyph = ">"
	case ">":
		g.guard.Glyph = "v"
	case "v":
		g.guard.Glyph = "<"
	case "<":
		g.guard.Glyph = "^"
	}
}

func (g *Grid) Step() error {
	switch g.guard.Glyph {
	case "^":
		if g.guard.Y == 0 {
			return ErrOutOfBounds
		}
		if g.rawMap[g.guard.Y-1][g.guard.X] == "#" {
			return ErrOutOfBounds
		}
		g.trackedMap[g.guard.Y][g.guard.X] = "X"
		g.guard.X -= stepUp.Xstep
		g.guard.Y -= stepUp.Ystep
		g.trackedMap[g.guard.Y][g.guard.X] = g.guard.Glyph
		return nil
	case "v":
		if g.guard.Y == len(g.rawMap) {
			return ErrOutOfBounds
		}
	case ">":
		if g.guard.X == len(g.rawMap[0])-1 {
			return ErrOutOfBounds
		}
	case "<":
		if g.guard.X == 0 {
			return ErrOutOfBounds
		}
	}
	panic("YIKES")
}

type Grid struct {
	rawMap     [][]string
	trackedMap [][]string
	guard      Guard
}

func ReadInput(path string) (*Grid, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	grid := [][]string{}
	g := Guard{
		Glyph:     "^",
		Direction: stepUp,
		X:         0,
		Y:         0,
	}
	y := 0
	for scanner.Scan() {
		err := scanner.Err()
		if err != nil {
			return nil, err
		}
		entries := strings.Split(scanner.Text(), "")
		for x, char := range entries {
			if char == "^" {
				g.X = x
				g.Y = y
			}
		}
		grid = append(grid, entries)
		y++

	}
	dupMap := make([][]string, len(grid))
	for i := range grid {
		dupMap[i] = make([]string, len(grid[i]))
		copy(dupMap[i], grid[i])
	}
	return &Grid{
		rawMap:     grid,
		trackedMap: dupMap,
		guard:      g,
	}, nil
}

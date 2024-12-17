package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func ReadFileBytes(path string) []byte {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return fileBytes
}

func ReadFileLines(path string) []string {
	fileBytes := ReadFileBytes(path)
	lines := []string{}
	separator := []byte("\n")
	for _, line := range bytes.Split(fileBytes, separator) {
		if string(line) != "" {
			lines = append(lines, string(line))
		}
	}
	return lines
}

func ConstructGrid(lines []string) [][]string {
	grid := [][]string{}
	for _, line := range lines {
		row := []string{}
		for _, char := range strings.Split(line, "") {
			row = append(row, char)
		}
		grid = append(grid, row)
	}
	fmt.Println(grid)
	return grid
}

var directions [][]int = [][]int{
	{0, 1},   // up
	{0, -1},  //down
	{1, 0},   // right
	{-1, 0},  // left
	{1, 1},   // up right
	{-1, 1},  // up left
	{1, -1},  // down right
	{-1, -1}, // down left
}

var wordList []string = []string{"X", "M", "A", "S"}

func FindXMAS(x, y, wordPosition int, direction []int, grid [][]string) bool {
	xNext := x + direction[0]
	yNext := y + direction[1]
	if wordPosition > len(wordList)-1 {
		return true
	}

	if xNext < 0 || xNext >= len(grid) || yNext < 0 || yNext >= len(grid[x]) {
		return false
	}

	if grid[xNext][yNext] == wordList[wordPosition] {
		return FindXMAS(xNext, yNext, wordPosition+1, direction, grid)
	}
	return false

}

func FindMAS(x, y int, grid [][]string) bool {
	xL := x - 1
	yT := y + 1
	xR := x + 1
	yD := y - 1

	// 0 3 2 1
	if xL < 0 || xL >= len(grid) || xR < 0 || xR >= len(grid) || yT < 0 || yT >= len(grid[xL]) || yT >= len(grid[xR]) || yD < 0 || yD >= len(grid[xL]) || yD >= len(grid[xR]) {
		return false
	}
	if ((grid[xL][yT] == wordList[1] && grid[xR][yD] == wordList[3]) || (grid[xL][yT] == wordList[3] && grid[xR][yD] == wordList[1])) && ((grid[xR][yT] == wordList[1] && grid[xL][yD] == wordList[3]) || (grid[xR][yT] == wordList[3] && grid[xL][yD] == wordList[1])) {
		return true
	}

	return false

}

func TraverseGrid2(grid [][]string) int {
	score := 0
	for x, row := range grid {
		for y, char := range row {
			if char == wordList[2] {
				if FindMAS(x, y, grid) {
					score += 1
				}

			}
		}
	}
	return score
}

func TraverseGrid(grid [][]string) int {
	score := 0
	for x, row := range grid {
		for y, char := range row {
			if char == wordList[0] {
				// traverse all directions and check if we find xmas
				for _, direction := range directions {
					if FindXMAS(x, y, 1, direction, grid) {
						score += 1
					}
				}
			}
		}
	}
	return score
}

func main() {
	lines := ReadFileLines("./input.txt")
	grid := ConstructGrid(lines)
	// score := TraverseGrid(grid)
	score := TraverseGrid2(grid)
	fmt.Println(score)
}

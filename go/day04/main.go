package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	crossInput := ReadInput("./test_input.txt")
	crossInput.checkY(1, 1)
	for _, entry := range crossInput.grid {
		fmt.Println(entry)
	}
}

func (c *Crossifier) checkY(x, y int) {
	start := c.grid[x][y]
	yLen := len(c.grid[x])
	xLen := len(c.grid[x][0])
	fmt.Println(yLen, xLen, start)
}

func (c *Crossifier) checkX(x, y int) {
}

func (c *Crossifier) heckDiag(x, y int) {

}

type Crossifier struct {
	grid [][]string
}

func ReadInput(path string) Crossifier {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	lines := [][]string{}
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		lines = append(lines, strings.Split(scanner.Text(), ""))
	}

	return Crossifier{grid: lines}
}

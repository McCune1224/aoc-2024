package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
)

func main() {
	fmt.Printf("PART1: %d\n", Part1())
	fmt.Printf("PART2: %d\n", Part2())
}
func Part1() int {

	left, right := ParseInput(ReadInput("day01.txt"))

	slices.Sort(left)
	slices.Sort(right)
	sum := 0
	for i := range left {
		sum += GetDistance(left[i], right[i])
	}
	return sum
}

func GetDistance(left, right int) int {
	return int(math.Abs(float64(left - right)))
}

func ReadInput(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func ParseInput(list []string) (left []int, right []int) {
	left = []int{}
	right = []int{}
	for _, entry := range list {
		nonWhitespace := true
		i := 0
		lEntry := ""
		for nonWhitespace {
			char := string(entry[i])
			if char == " " {
				nonWhitespace = false
				continue
			}
			lEntry += char
			i++
		}
		rEntry := entry[i+3:]
		lInt, err := strconv.Atoi(lEntry)
		if err != nil {
			log.Fatal(err)
		}
		rInt, err := strconv.Atoi(rEntry)
		if err != nil {
			log.Fatal(err)
		}

		left = append(left, lInt)
		right = append(right, rInt)

	}
	return left, right
}

func CalculateSimilarityScore(left []int, right []int) int {
	rightSet := make(map[int]int)
	for i := range right {
		_, rightOk := rightSet[right[i]]
		if rightOk {
			rightSet[right[i]] += 1
		} else {
			rightSet[right[i]] = 1
		}
	}

	score := 0
	for _, entry := range left {
		val, ok := rightSet[entry]
		if !ok {
			continue
		}
		score += val * entry
	}
	return score
}

func Part2() int {
	left, right := ParseInput(ReadInput("day01.txt"))
	return CalculateSimilarityScore(left, right)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Part01())
	fmt.Println(Part02())
}

func Part01() int {
	records := readInput("input.txt")

	// 1. The levels are either all increasing or all decresasing
	// 2. Any two adjacent levels differ by at least one and at most three
	safeTally := 0
	for _, record := range records {
		// fmt.Println(checkAdjacent(record))
		// fmt.Println(isAscOrDesc(record))
		if checkAdjacent(record) == true && isAscOrDesc(record) == true {
			safeTally += 1
			continue
		}
	}
	return safeTally
}

func Part02() int {
	records := readInput("input.txt")

	// 1. The levels are either all increasing or all decresasing
	// 2. Any two adjacent levels differ by at least one and at most three
	safeTally := 0
	for _, record := range records {
		if checkAdjacent(record) == true && isAscOrDesc(record) == true {
			safeTally += 1
			continue
		} else {
			permutations := generateCombinations(record)
			for _, perm := range permutations {
				if checkAdjacent(perm) == true && isAscOrDesc(perm) == true {
					safeTally += 1
					break
				}
			}
		}
	}
	return safeTally
}

func generateCombinations(numbers []int) [][]int {
	result := [][]int{}

	for i := 0; i < len(numbers); i++ {
		combination := make([]int, 0, len(numbers)-1)
		combination = append(combination, numbers[:i]...)
		combination = append(combination, numbers[i+1:]...)
		result = append(result, combination)
	}

	return result
}

func isAscOrDesc(record []int) bool {
	if record[0]-record[1] == 0 {
		return false
	}
	if record[0]-record[1] < 0 {
		for i := 0; i < len(record)-1; i++ {
			diff := record[i] - record[i+1]
			if diff > 0 || diff == 0 {
				return false
			}
		}
		return true
	} else {
		for i := 0; i < len(record)-1; i++ {
			diff := record[i] - record[i+1]
			if diff < 0 || diff == 0 {
				return false
			}
		}
		return true
	}

}
func checkAdjacent(record []int) bool {
	for i := range record {
		// fmt.Printf("----------- %d -----------\n", i)
		if i == 0 {
			diff := absDiffInt(record[0], record[1])
			// fmt.Println(diff, record[0], record[1])
			if diff < 1 || diff > 3 {
				return false
			}
			continue
		}
		if i == len(record)-1 {
			diff := absDiffInt(record[i], record[i-1])
			// fmt.Println(diff, record[i], record[i-1])
			if diff < 1 || diff > 3 {
				return false
			}
			continue
		}
		lDiff := absDiffInt(record[i], record[i-1])
		rDiff := absDiffInt(record[i], record[i+1])
		if lDiff < 1 || lDiff > 3 {
			return false
		}
		if rDiff < 1 || rDiff > 3 {
			return false
		}

	}
	return true
}

func absDiffInt(x int, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func readInput(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	records := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		strNums := strings.Split(line, " ")
		record := []int{}
		for _, str := range strNums {
			level, err := strconv.Atoi(str)
			if err != nil {
				log.Fatal(err)
			}
			record = append(record, level)
		}
		records = append(records, record)
	}

	return records
}

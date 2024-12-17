package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type OrderingRule struct {
	X int
	Y int
}

func main() {
	Part2()
}

func Part2() {
	// rules, updates, err := ReadInput("./test_input.txt")
	rules, updates, err := ReadInput("./test_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(rules)

	grX := groupRulesByX(rules)

	invalidUpdates := make([][]int, 0)
	for _, update := range updates {
		validUpdate := isValidUpdateLine(update, grX)
		if !validUpdate {
			invalidUpdates = append(invalidUpdates, update)
		}
	}

	tally := 0
	for _, validUpdate := range invalidUpdates {
		correctedUpdate := CorrectOrdering(validUpdate, rules)
		tally += correctedUpdate[len(validUpdate)/2]
	}
	fmt.Println(tally)
}

func CorrectOrdering(update []int, rules []OrderingRule) []int {
	prevItems := []int{update[0]}
	grX := groupRulesByX(rules)
	for i := 1; i < len(update)-1; i++ {
		item := update[i]
		if existingRules, ok := grX[item]; ok {
			for _, v := range existingRules {
				if slices.Contains(prevItems, v) {
					idx := slices.Index(update, v)
					fmt.Printf("%d MUST come before %d at index %d\n", item, update[idx], idx)
					fmt.Println(update)
					fmt.Println(item, grX[item])
				}
			}
		}
	}
	return update
}

func isValidUpdateLine(update []int, grX map[int][]int) bool {
	prevItems := []int{}
	for _, item := range update {
		if existingRules, ok := grX[item]; ok {
			for _, v := range existingRules {
				if slices.Contains(prevItems, v) {
					return false
				}
			}
		}
		prevItems = append(prevItems, item)
	}
	return true
}

func groupRulesByX(rules []OrderingRule) map[int][]int {
	ruleMap := make(map[int][]int)
	for _, rule := range rules {
		currentRules, ok := ruleMap[rule.X]
		if ok {
			ruleMap[rule.X] = append(currentRules, rule.Y)
		} else {
			ruleMap[rule.X] = []int{rule.Y}
		}
	}
	return ruleMap
}

func groupRulesByY(rules []OrderingRule) map[int][]int {
	ruleMap := make(map[int][]int)
	for _, rule := range rules {
		currentRules, ok := ruleMap[rule.Y]
		if ok {
			ruleMap[rule.Y] = append(currentRules, rule.X)
		} else {
			ruleMap[rule.Y] = []int{rule.X}
		}
	}
	return ruleMap
}

func ReadInput(path string) ([]OrderingRule, [][]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}

	scanner := bufio.NewScanner(file)

	part2 := false
	var rules []OrderingRule
	var updates [][]int

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Println("Failed to scan input:", err)
			return nil, nil, err
		}
		line := scanner.Text()
		if line == "" {
			part2 = true
			continue
		}
		if !part2 {
			lineSplit := strings.Split(line, "|")
			x, _ := strconv.Atoi(lineSplit[0])
			y, _ := strconv.Atoi(lineSplit[1])
			rules = append(rules, OrderingRule{X: x, Y: y})
			continue
		}
		rulesSplit := strings.Split(line, ",")
		intRuleSplit := make([]int, 0)
		for _, rule := range rulesSplit {
			num, _ := strconv.Atoi(rule)
			intRuleSplit = append(intRuleSplit, num)
		}
		updates = append(updates, intRuleSplit)
	}
	return rules, updates, nil
}

func Part1() {
	// rules, updates, err := ReadInput("./test_input.txt")
	rules, updates, err := ReadInput("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(rules)

	grX := groupRulesByX(rules)

	validUpdates := make([][]int, 0)
	for _, update := range updates {
		validUpdate := isValidUpdateLine(update, grX)
		if validUpdate {
			validUpdates = append(validUpdates, update)
		}
	}
	tally := 0
	for _, validUpdate := range validUpdates {
		tally += validUpdate[len(validUpdate)/2]
	}
	fmt.Println(tally)
}

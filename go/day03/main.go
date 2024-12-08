package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))
func main() {
	foo := LoadInput("./input.txt")
	foo.ClipMultStatements()
	// fmt.Println(foo.statements)
	// fmt.Println(foo.SolveMultStatements())
	fmt.Println(foo.SolveMultStatements())
	// // entry := strings.Split(b, "don't()")
	// entry := strings.Split(b, "do()")
	// for _, e := range entry {
	// 	fmt.Println(e)
	// }
}

type Parser struct {
	buff       string
	statements []Statement
}

type Statement struct {
	mult   string
	Active bool
}

// Really janky, but just splitting a new parser off when we find a 'do()' or 'dont()' line, discarding the don't lines
// func (p *Parser) SplitParser() []Parser {
// }

func (p *Parser) SolveMultStatements() int {
	sum := 0
	for _, statement := range p.statements {
		if statement.Active {
			start := strings.Index(statement.mult, "(")
			end := strings.Index(statement.mult, ")")
			numbers := strings.Split(statement.mult[start+1:end], ",")
			left, _ := strconv.Atoi(numbers[0])
			right, _ := strconv.Atoi(numbers[1])
			sum += left * right
		}
	}
	return sum
}

func (p *Parser) ClipMultStatements() []Statement {
	i := 0
	multStatements := []Statement{}
	active := true
	for i != len(p.buff)-4 {
		if i > len(p.buff)-8 {
			break
		}
		if p.buff[i] == 'd' {
			if strings.Contains(p.buff[i:i+7], "do()") {
				active = true
			}
			if strings.Contains(p.buff[i:i+7], "don't()") {
				active = false
			}
		}
		if p.buff[i:i+4] == "mul(" {
			startIdx := i
			i += 3

			numTally := 0
			isDigitCheck := isDigit(rune(p.buff[i+1]))
			for isDigitCheck {
				if numTally > 3 {
					i = startIdx
				}
				i++
				numTally++
				isDigitCheck = isDigit(rune(p.buff[i+1]))
				// fmt.Println(p.buff[startIdx:i])
			}
			if rune(p.buff[i+1]) == ',' {
				i++
			} else {
				continue
			}
			numTally = 0
			isDigitCheck = isDigit(rune(p.buff[i+1]))
			for isDigitCheck {
				if numTally > 3 {
					i = startIdx
				}
				i++
				numTally++
				isDigitCheck = isDigit(rune(p.buff[i+1]))
			}
			if rune(p.buff[i+1]) == ')' {
				i++
				multStatements = append(multStatements, Statement{mult: p.buff[startIdx : i+1], Active: active})
			}
		}
		i++
	}
	p.statements = multStatements
	return multStatements
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func LoadInput(path string) *Parser {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	text := ""
	for scanner.Scan() {
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}
		text += scanner.Text()
	}
	return &Parser{buff: text}
}

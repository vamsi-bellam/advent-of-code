package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func isDigit(v rune) bool {
	return (v == '0' || v == '1' || v == '2' || v == '3' || v == '4' ||
		v == '5' || v == '6' || v == '7' || v == '8' || v == '9')
}

func findSum(considerConditonals bool) int {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	lastChar, lastLastChar := 'x', 'x'
	num1, num2 := "", ""
	sum := 0
	isCommaPassed, isMulEnabled := false, true

	for scanner.Scan() {
		line := scanner.Text()

		for _, v := range line {

			if considerConditonals &&
				v == 'd' || (v == 'o' && lastChar == 'd') ||
				(v == 'n' && lastChar == 'o') ||
				(v == '\'' && lastChar == 'n') ||
				(v == 't' && lastChar == '\'') ||
				(v == '(' && (lastChar == 'o' || lastChar == 't')) {
				lastLastChar = lastChar
				lastChar = v
			} else if considerConditonals && v == ')' && lastChar == '(' {
				if lastLastChar == 't' {
					isMulEnabled = false
				} else if lastLastChar == 'o' {
					isMulEnabled = true
				}
				lastLastChar = 'x'
			} else if isMulEnabled {
				if v == 'm' ||
					(v == 'u' && lastChar == 'm') ||
					(v == 'l' && lastChar == 'u') ||
					(v == '(' && lastChar == 'l') {
					lastChar = v
				} else if isDigit(v) &&
					(lastChar == '(' || (!isCommaPassed && isDigit(lastChar))) {
					lastChar = v
					num1 += string(v)
				} else if v == ',' && isDigit(lastChar) {
					lastChar = v
					isCommaPassed = true
				} else if isDigit(v) &&
					(lastChar == ',' || (isCommaPassed && isDigit(lastChar))) {
					lastChar = v
					num2 += string(v)
				} else if v == ')' && (isCommaPassed && isDigit(lastChar)) {
					v1, err1 := strconv.Atoi(num1)
					if err1 != nil {
						log.Fatalf("Unable to parse num1 to integer %v", num1)
					}
					v2, err2 := strconv.Atoi(num2)
					if err2 != nil {
						log.Fatalf("Unable to parse num2 to integer %v", num2)
					}
					sum += (v1 * v2)
					lastChar = 'x'
					num1 = ""
					num2 = ""
					isCommaPassed = false
				} else {
					lastChar = 'x'
					num1 = ""
					num2 = ""
					isCommaPassed = false
				}
			}

		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()

	return sum
}

func main() {

	fmt.Printf("Sum without considering conditionals %d \n", findSum(false))

	fmt.Printf("Sum with considering conditionals %d \n", findSum(true))
}

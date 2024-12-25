package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func conNums(num1, num2 int) int {
	num2cpy := num2
	for num2 > 0 {
		num1 *= 10
		num2 /= 10
	}
	return num1 + num2cpy
}

func isValid(target int, sum int, sl []int, enableCons bool) bool {
	if len(sl) == 0 {
		return target == sum
	}
	first, rest := sl[0], sl[1:]
	mulSum := sum
	if sum == 0 {
		mulSum = 1
	}
	return isValid(target, sum+first, rest, enableCons) ||
		isValid(target, mulSum*first, rest, enableCons) ||
		(enableCons && isValid(target, conNums(sum, first), rest, enableCons))
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	equations := make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, ":")

		restList := strings.Split(strings.Trim(nums[1], " "), " ")

		values := make([]int, len(restList)+1)

		target, err := strconv.Atoi(nums[0])
		if err != nil {
			log.Fatal(err)
		}
		values[0] = target

		for i, v := range restList {
			val, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			values[i+1] = val
		}

		equations = append(equations, values)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()

	sum := 0
	consSum := 0

	for _, equation := range equations {
		if isValid(equation[0], 0, equation[1:], false) {
			sum += equation[0]
		}

		if isValid(equation[0], 0, equation[1:], true) {
			consSum += equation[0]
		}

	}

	fmt.Printf("Sum of satisfied targets : %d \n", sum)
	fmt.Printf("Sum of satisfied targets with cons : %d \n", consSum)
}

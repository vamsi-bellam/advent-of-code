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

func getDigits(num int) []int {
	digits := make([]int, 0)

	if num == 0 {
		digits = append(digits, 0)
	}

	for num > 0 {
		rem := num % 10
		digits = append(digits, rem)
		num = num / 10
	}

	slices.Reverse(digits)
	return digits
}

func pow(a, b int) int {
	if b == 0 {
		return 1
	}
	return a * pow(a, b-1)
}

func makeNumber(digits []int) int {
	num := 0
	digitsLen := len(digits)
	for i, v := range digits {
		num += (pow(10, digitsLen-i-1) * v)
	}
	return num
}

func noOfStonesAfterBlinks(stones []int, blinks int) int {
	newStones := stones

	for blinks > 0 {

		stones = newStones
		newStones = make([]int, 0)

		for i, v := range stones {
			if v == 0 {
				newStones = append(newStones, 1)
			} else if digits := getDigits(v); len(digits)%2 == 0 {
				newStones = append(newStones, makeNumber(digits[:len(digits)/2]))
				newStones = append(newStones, makeNumber(digits[len(digits)/2:]))
			} else {
				newStones = append(newStones, stones[i]*2024)
			}
		}
		blinks--
	}
	return len(newStones)
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	stones := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		numsInStr := strings.Split(line, " ")

		for _, v := range numsInStr {
			num, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			stones = append(stones, num)
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()

	stonesCopy := make([]int, len(stones))
	copy(stonesCopy, stones)

	fmt.Printf("stones after blinking 25 times - %d\n", noOfStonesAfterBlinks(stones, 25))
	fmt.Printf("stones after blinking 75 times - %d\n", noOfStonesAfterBlinks(stonesCopy, 75))
}

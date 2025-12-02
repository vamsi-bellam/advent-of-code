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

type BlinkAndStone struct {
	blink, stone int
}

func stonesAfterBlinks(stone int, blinks int, cache map[BlinkAndStone]int) int {

	if v, ok := cache[BlinkAndStone{blinks, stone}]; ok {
		return v
	}

	newStones := make([]int, 0)

	if stone == 0 {
		newStones = append(newStones, 1)
	} else if digits := getDigits(stone); len(digits)%2 == 0 {
		newStones = append(newStones, makeNumber(digits[:len(digits)/2]))
		newStones = append(newStones, makeNumber(digits[len(digits)/2:]))
	} else {
		newStones = append(newStones, stone*2024)
	}

	if blinks == 1 {
		count := len(newStones)
		cache[BlinkAndStone{blinks, stone}] = count
		return count
	}

	sum := 0

	for _, stone := range newStones {
		count := stonesAfterBlinks(stone, blinks-1, cache)
		cache[BlinkAndStone{blinks - 1, stone}] = count
		sum += count
	}

	cache[BlinkAndStone{blinks, stone}] = sum

	return sum
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

	sum := 0

	cache := make(map[BlinkAndStone]int)

	for _, stone := range stones {
		sum += stonesAfterBlinks(stone, 25, cache)
	}

	sum2 := 0

	for _, stone := range stones {
		sum2 += stonesAfterBlinks(stone, 75, cache)
	}

	fmt.Printf("stones after blinking 25 times - %d\n", sum)
	fmt.Printf("stones after blinking 75 times - %d\n", sum2)
}

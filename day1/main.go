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

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	leftArray := make([]int, 0)
	rightArray := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, "   ")

		firstNumber, err := strconv.Atoi(numbers[0])
		if err != nil {
			log.Fatal(err)
		}
		leftArray = append(leftArray, firstNumber)

		secondNumber, err := strconv.Atoi(numbers[1])
		if err != nil {
			log.Fatal(err)
		}
		rightArray = append(rightArray, secondNumber)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()

	slices.Sort(leftArray)

	slices.Sort(rightArray)

	totalDistance := 0

	for i := range leftArray {
		first, second := leftArray[i], rightArray[i]

		if first > second {
			totalDistance += first - second
		} else {
			totalDistance += second - first
		}
	}

	fmt.Println(totalDistance)

	freq := make(map[int]int)

	for _, value := range rightArray {
		freq[value] += 1
	}

	similarityScore := 0

	for _, value := range leftArray {
		count, ok := freq[value]
		if ok {
			similarityScore += value * count
		}
	}

	fmt.Println(similarityScore)
}

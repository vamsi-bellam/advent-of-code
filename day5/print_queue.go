package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Map(ls []string, f func(string) (int, error)) ([]int, error) {
	newLs := make([]int, len(ls))
	for i, v := range ls {
		val, err := f(v)
		if err != nil {
			return nil, err
		}
		newLs[i] = val
	}
	return newLs, nil
}
func isInSlice[T comparable](ls []T, v T) bool {
	for _, ele := range ls {
		if ele == v {
			return true
		}
	}
	return false
}
func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	logBook := make(map[int][]int, 0)
	updates := make([][]int, 0)

	isRulesEnded := false
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			isRulesEnded = true
		} else if !isRulesEnded {
			nums := strings.Split(line, "|")
			v1, err1 := strconv.Atoi(nums[0])
			if err1 != nil {
				log.Fatalf("Unable to parse num1 to integer %v", nums[0])
			}
			v2, err2 := strconv.Atoi(nums[1])
			if err2 != nil {
				log.Fatalf("Unable to parse num2 to integer %v", nums[1])
			}

			/*
				Maintain a log book such that elements that should come
				before an element are stored in a list
				{2 : [8, 9, 10]} here 8, 9, 10 should come before 2
			*/

			val, ok := logBook[v2]
			if !ok {

				logBook[v2] = []int{v1}
			} else {
				val := append(val, v1)
				logBook[v2] = val
			}
		} else {
			recs := strings.Split(line, ",")
			intOfRecs, err := Map(recs, strconv.Atoi)

			if err != nil {
				log.Fatal(err)
			}
			updates = append(updates, intOfRecs)
		}
	}
	file.Close()

	sum := 0
	fixedSum := 0

	for _, update := range updates {
		isValid := true
		len := len(update)
	out:
		for i := 0; i < len; i++ {
			beforeList, ok := logBook[update[i]]
			if ok {
				for j := i + 1; j < len; j++ {
					// if an element2 is before list of element1 then it is
					// not valid update
					if isInSlice(beforeList, update[j]) {
						isValid = false
						break out
					}
				}
			}
		}

		if isValid {
			sum += update[len/2]
		} else {
			// If not valid, sort the update such that it follows rule in logBook
			sort.Slice(update, func(i, j int) bool {
				if _, ok := logBook[update[i]]; ok {
					for _, val := range logBook[update[i]] {
						if val == update[j] {
							return false
						}
					}
				}
				return true
			})
			fixedSum += update[len/2]
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Sum is %d \n", sum)
	fmt.Printf("Sum After Fixing Only Incorrect Updates %d \n", fixedSum)
}

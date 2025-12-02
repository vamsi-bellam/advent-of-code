package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

func visitIfPossible(hoofMap [][]int, i, j, maxI, maxJ int, lastLevel int,
	visited map[Point]int) {

	if i >= 0 && i < maxI && j >= 0 && j < maxJ {
		currLevel := hoofMap[i][j]
		if currLevel == lastLevel+1 {
			if currLevel == 9 {
				if _, ok := visited[Point{i, j}]; !ok {
					visited[Point{i, j}] = 1
				} else {
					visited[Point{i, j}] += 1
				}
			}
			visitIfPossible(hoofMap, i-1, j, maxI, maxJ, currLevel, visited)
			visitIfPossible(hoofMap, i+1, j, maxI, maxJ, currLevel, visited)
			visitIfPossible(hoofMap, i, j-1, maxI, maxJ, currLevel, visited)
			visitIfPossible(hoofMap, i, j+1, maxI, maxJ, currLevel, visited)
		}
	}

}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	hoofMap := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		numsInStr := strings.Split(line, "")
		nums := make([]int, len(numsInStr))
		for i, v := range numsInStr {
			num, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			nums[i] = num
		}

		hoofMap = append(hoofMap, nums)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()

	sum := 0
	sum2 := 0

	rows, rowLen := len(hoofMap), len(hoofMap[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < rowLen; j++ {
			if hoofMap[i][j] == 0 {
				// track visited 9's

				visited := make(map[Point]int)
				visitIfPossible(hoofMap, i, j, rows, rowLen, -1, visited)

				// For part - 1 : score is nothing but all visited 9's
				sum += len(visited)

				// For part - 2: score is nothing but all different ways to visit 9's
				for _, v := range visited {
					sum2 += v
				}
			}
		}
	}
	fmt.Printf("Sum of the scores - Part - 1 : %d \n", sum)
	fmt.Printf("Sum of the scores - Part - 2 : %d \n", sum2)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func findXMAS(data [][]string) int {

	count := 0

	cols := len(data)
	rowLen := len(data[0])

	for i, iv := range data {
		for j, jv := range iv {
			// Horizantal Right
			if jv == "X" && j+3 < rowLen &&
				iv[j+1] == "M" && iv[j+2] == "A" && iv[j+3] == "S" {
				count++

			}

			// Horizantal Left ~ Reverse
			if jv == "X" && j-3 >= 0 &&
				iv[j-1] == "M" && iv[j-2] == "A" && iv[j-3] == "S" {
				count++

			}

			// Vertical Down
			if jv == "X" && i+3 < cols &&
				data[i+1][j] == "M" && data[i+2][j] == "A" &&
				data[i+3][j] == "S" {
				count++

			}

			// Vertical Up ~ Reverse
			if jv == "X" && i-3 >= 0 &&
				data[i-1][j] == "M" && data[i-2][j] == "A" &&
				data[i-3][j] == "S" {
				count++

			}

			// Diagonal 1 (i, j), (i-1, j-1), (i-2 , j-2), (i-3, j-3)

			if jv == "X" && i-3 >= 0 && j-3 >= 0 &&
				data[i-1][j-1] == "M" && data[i-2][j-2] == "A" &&
				data[i-3][j-3] == "S" {
				count++

			}

			// Diagonal (i, j), (i-1, j+1), (i-2 , j+2), (i-3, j+3)

			if jv == "X" && i-3 >= 0 && j+3 < rowLen &&
				data[i-1][j+1] == "M" && data[i-2][j+2] == "A" &&
				data[i-3][j+3] == "S" {
				count++

			}

			// Diagonal 4 (i, j), (i+1, j-1), (i+2 , j-2), (i+3, j-3)

			if jv == "X" && i+3 < cols && j-3 >= 0 &&
				data[i+1][j-1] == "M" && data[i+2][j-2] == "A" &&
				data[i+3][j-3] == "S" {
				count++

			}

			// Diagonal 3 (i, j), (i+1, j+1), (i+2 , j+2), (i+3, j+3)
			if jv == "X" && i+3 < cols && j+3 < rowLen &&
				data[i+1][j+1] == "M" && data[i+2][j+2] == "A" &&
				data[i+3][j+3] == "S" {
				count++

			}

		}
	}
	return count
}

func findXShapedMAS(data [][]string) int {
	count := 0

	cols := len(data)
	rowLen := len(data[0])

	for i := 0; i+2 < cols; i++ {
		for j := 0; j+2 < rowLen; j++ {
			if data[i+1][j+1] == "A" &&
				((data[i][j] == "M" && data[i+2][j+2] == "S") ||
					(data[i][j] == "S" && data[i+2][j+2] == "M")) &&
				((data[i][j+2] == "M" && data[i+2][j] == "S") ||
					(data[i][j+2] == "S" && data[i+2][j] == "M")) {
				count++
			}
		}
	}
	return count
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	data := make([][]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		letters := strings.Split(line, "")
		data = append(data, letters)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()

	fmt.Printf("Total count of word XMAS - %d \n", findXMAS(data))
	fmt.Printf("Total count of word X-MAS - %d \n", findXShapedMAS(data))
}

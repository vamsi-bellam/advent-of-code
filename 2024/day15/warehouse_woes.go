package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Point struct {
	X, Y int
}

func (p *Point) Update(x, y int) {
	p.X = x
	p.Y = y
}

func findRobo(warehouse [][]string) Point {

	for i, row := range warehouse {
		for j, v := range row {
			if v == "@" {

				return Point{i, j}
			}
		}
	}
	return Point{0, 0}
}
func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	warehouse := make([][]string, 0)
	movements := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		s := strings.Split(line, "")

		if s[0] == "#" {
			warehouse = append(warehouse, s)
		} else {
			movements = append(movements, s...)
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()

	rows := len(warehouse)
	rowLen := len(warehouse[0])

	robot := findRobo(warehouse)

	for _, m := range movements {
		i, j := robot.X, robot.Y

		if m == "^" {
			next := 0
			for x := i - 1; x >= 0; x-- {

				if warehouse[x][j] == "." {
					next = x
					break
				} else if warehouse[x][j] == "#" {
					break
				}
			}
			if next != 0 {
				for x := next + 1; x <= i; x++ {
					warehouse[x-1][j] = warehouse[x][j]
				}
				warehouse[i][j] = "."
				robot.Update(i-1, j)
			}

		} else if m == ">" {
			next := rowLen
			for x := j + 1; x < rowLen; x++ {

				if warehouse[i][x] == "." {
					next = x
					break
				} else if warehouse[i][x] == "#" {
					break
				}
			}
			if next != rowLen {
				fmt.Println(next)
				for x := next - 1; x >= j; x-- {
					warehouse[i][x+1] = warehouse[i][x]
				}
				warehouse[i][j] = "."
				robot.Update(i, j+1)
			}
		} else if m == "v" {
			next := rows
			for x := i + 1; x < rows; x++ {

				if warehouse[x][j] == "." {
					next = x
					break
				} else if warehouse[x][j] == "#" {
					break
				}
			}
			if next != rows {
				for x := next - 1; x >= i; x-- {
					warehouse[x+1][j] = warehouse[x][j]
				}
				warehouse[i][j] = "."
				robot.Update(i+1, j)
			}
		} else if m == "<" {
			next := 0
			for x := j - 1; x >= 0; x-- {

				if warehouse[i][x] == "." {
					next = x
					break
				} else if warehouse[i][x] == "#" {
					break
				}
			}
			if next != 0 {
				for x := next + 1; x <= j; x++ {
					warehouse[i][x-1] = warehouse[i][x]
				}
				warehouse[i][j] = "."
				robot.Update(i, j-1)
			}
		}

		// fmt.Printf("After move %s \n", m)
		// for _, v := range warehouse {
		// 	for _, g := range v {
		// 		fmt.Print(g)
		// 	}
		// 	fmt.Println()
		// }
	}

	sum := 0

	for i, row := range warehouse {
		for j, v := range row {
			if v == "O" {
				sum += 100*i + j

			}
		}
	}

	fmt.Printf("Sum for part - 1 : %d \n", sum)

	newWarehouse := make([][]string, rows)

	for i, row := range warehouse {
		newRow := make([]string, rowLen*2)
		newWarehouse = append(newWarehouse, newRow)
		for j, v := range row {
			if v == "#" {
				newWarehouse[i][j*2] = "#"
				newWarehouse[i][j*2+1] = "#"
			} else if v == "O" {
				newWarehouse[i][j*2] = "["
				newWarehouse[i][j*2+1] = "]"

			} else if v == "." {
				newWarehouse[i][j*2] = "."
				newWarehouse[i][j*2+1] = "."
			} else if v == "@" {
				newWarehouse[i][j*2] = "@"
				newWarehouse[i][j*2+1] = "."
			}
		}
	}

}

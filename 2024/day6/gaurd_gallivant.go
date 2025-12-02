package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func isBorder(i, j, iLen, jLen int, dir string) bool {
	if dir == "up" {
		return i == 0
	} else if dir == "down" {
		return i == iLen-1
	} else if dir == "left" {
		return j == 0
	} else if dir == "right" {
		return j == jLen-1
	}
	return false
}

func findGuradPos(guardMap [][]string) (int, int, string, bool) {
	for i, iv := range guardMap {
		for j, jv := range iv {
			if jv == "^" {
				return i, j, "up", true
			} else if jv == ">" {
				return i, j, "right", true
			} else if jv == "<" {
				return i, j, "left", true
			} else if jv == "v" {
				return i, j, "down", true
			}
		}
	}
	return 0, 0, "", false
}

func findUniqPos(guardMap [][]string) (int, bool) {
	i, j, dir, ok := findGuradPos(guardMap)
	loop := false

	logBook := make(map[string]bool)

	if !ok {
		log.Fatalf("Guard not found in map!")
	}

	iLen, jLen := len(guardMap), len(guardMap[0])

	uniqPos := 1

	guardMap[i][j] = "X"

	for {
		// is a loop if same i, j with same direction
		key := string(i) + string(j) + dir
		already, ok := logBook[key]

		if ok && already {
			loop = true
			break
		}

		logBook[key] = true

		if isBorder(i, j, iLen, jLen, dir) {
			break
		}

		if dir == "up" {
			if guardMap[i-1][j] != "#" {
				i = i - 1
				if guardMap[i][j] != "X" {
					uniqPos++
					guardMap[i][j] = "X"
				}
			} else {
				dir = "right"
			}
		} else if dir == "down" {
			if guardMap[i+1][j] != "#" {
				i = i + 1
				if guardMap[i][j] != "X" {
					uniqPos++
					guardMap[i][j] = "X"
				}
			} else {
				dir = "left"
			}

		} else if dir == "left" {
			if guardMap[i][j-1] != "#" {
				j = j - 1
				if guardMap[i][j] != "X" {
					uniqPos++
					guardMap[i][j] = "X"
				}
			} else {
				dir = "up"
			}
		} else if dir == "right" {
			if guardMap[i][j+1] != "#" {
				j = j + 1
				if guardMap[i][j] != "X" {
					uniqPos++
					guardMap[i][j] = "X"
				}
			} else {
				dir = "down"
			}
		}

	}
	return uniqPos, loop
}

func deepCopySlice[T comparable](sl [][]T) [][]T {
	slcpy := make([][]T, 0)
	for _, in := range sl {
		incpy := make([]T, len(in))
		copy(incpy, in)
		slcpy = append(slcpy, incpy)
	}
	return slcpy
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	guardMap := make([][]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		pos := strings.Split(line, "")
		guardMap = append(guardMap, pos)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()

	uniqPosCount, isLoop := findUniqPos(deepCopySlice(guardMap))

	if isLoop {
		log.Fatalf("Gurad stuck in loop!")
	}

	fmt.Printf("Unique positions guard can travel is %d \n", uniqPosCount)

	// Part - 2
	// Brute force way - keeping obstacle in very possible path
	obstPos := 0

	for i, iv := range guardMap {
		for j, jv := range iv {
			if jv == "." {
				guardMapCp := deepCopySlice(guardMap)
				guardMapCp[i][j] = "#"
				if _, isLoop := findUniqPos(guardMapCp); isLoop {
					obstPos++
				}
				guardMapCp[i][j] = "."
			}
		}
	}
	fmt.Printf("Obst count %d \n", obstPos)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Point struct {
	X int
	Y int
}

func isValidPoint(p Point, maxX, maxY int) bool {
	return (p.X >= 0 && p.X <= maxX) && (p.Y >= 0 && p.Y <= maxY)
}

func countAntiNodes(antennaMap [][]string, antennaPoints map[string][]Point,
	countAll bool) int {
	maxX, maxY := len(antennaMap[0])-1, len(antennaMap)-1
	uniqAntiNodes := make(map[Point]bool)
	count := 0

	for _, points := range antennaPoints {
		len := len(points)
		for i := 0; i < len; i++ {

			beforePoint := points[i]

			for j := i + 1; j < len; j++ {

				afterPoint := points[j]
				diffX := afterPoint.X - beforePoint.X
				diffY := afterPoint.Y - beforePoint.Y

				if countAll {
					antiNode1 := points[i]
					antiNode2 := points[j]

					for isValidPoint(antiNode1, maxX, maxY) {
						if _, ok := uniqAntiNodes[antiNode1]; !ok {
							uniqAntiNodes[antiNode1] = true
							count++
						}
						antiNode1 = Point{antiNode1.X - diffX, antiNode1.Y - diffY}
					}

					for isValidPoint(antiNode2, maxX, maxY) {
						if _, ok := uniqAntiNodes[antiNode2]; !ok {
							uniqAntiNodes[antiNode2] = true
							count++
						}
						antiNode2 = Point{antiNode2.X + diffX, antiNode2.Y + diffY}
					}

				} else {

					antiNode1 := Point{beforePoint.X - diffX, beforePoint.Y - diffY}
					antiNode2 := Point{afterPoint.X + diffX, afterPoint.Y + diffY}

					if isValidPoint(antiNode1, maxX, maxY) {
						if _, ok := uniqAntiNodes[antiNode1]; !ok {
							uniqAntiNodes[antiNode1] = true
							count++
						}
					}

					if isValidPoint(antiNode2, maxX, maxY) {
						if _, ok := uniqAntiNodes[antiNode2]; !ok {
							uniqAntiNodes[antiNode2] = true
							count++
						}

					}

				}
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

	antennaMap := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		antennaMap = append(antennaMap, strings.Split(line, ""))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()

	antennaPoints := make(map[string][]Point)

	for y, row := range antennaMap {
		for x, antenna := range row {
			if antenna != "." {
				v, ok := antennaPoints[antenna]
				if !ok {
					antennaPoints[antenna] = []Point{{x, y}}
				} else {
					antennaPoints[antenna] = append(v, Point{x, y})
				}
			}
		}
	}

	fmt.Printf("Unique antinode positions - %d\n",
		countAntiNodes(antennaMap, antennaPoints, false))

	fmt.Printf("Unique antinode positions including all - %d\n",
		countAntiNodes(antennaMap, antennaPoints, true))
}

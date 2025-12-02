package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func moveRobot(robot Robot, xMax, yMax, times int) Robot {

	X, Y := robot.Position.X, robot.Position.Y

	X = (X + robot.Velocity.X*times) % xMax
	Y = (Y + robot.Velocity.Y*times) % yMax

	if X < 0 {
		X += xMax
	} else if X >= xMax {
		X -= xMax
	}

	if Y < 0 {
		Y += yMax
	} else if Y >= yMax {
		Y -= yMax
	}

	return Robot{Pair{X, Y}, robot.Velocity}
}

type Pair struct {
	X, Y int
}

type Robot struct {
	Position, Velocity Pair
}

func getPair(s string) Pair {
	nums := strings.Split(strings.Split(s, "=")[1], ",")

	x, err := strconv.Atoi(nums[0])

	if err != nil {
		log.Fatal(err)
	}

	y, err := strconv.Atoi(nums[1])

	if err != nil {
		log.Fatal(err)
	}
	return Pair{x, y}
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	robots := make([]Robot, 0)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		pos, vel := getPair(line[0]), getPair(line[1])
		robots = append(robots, Robot{pos, vel})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()

	xMax, yMax := 101, 103
	q1, q2, q3, q4 := 0, 0, 0, 0

	for i := range robots {

		robots[i] = moveRobot(robots[i], xMax, yMax, 100)

		X, Y := robots[i].Position.X, robots[i].Position.Y

		if X < xMax/2 && Y < yMax/2 {
			q1 += 1
		}

		if X > xMax/2 && Y < yMax/2 {
			q2 += 1
		}

		if X < xMax/2 && Y > yMax/2 {
			q3 += 1
		}

		if X > xMax/2 && Y > yMax/2 {
			q4 += 1
		}
	}

	fmt.Printf("Mult is : %d \n", q1*q2*q3*q4)
}

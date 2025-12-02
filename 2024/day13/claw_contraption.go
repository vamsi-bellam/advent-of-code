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

type Game struct {
	ButtonA Point
	ButtonB Point
	Prize   Point
}

// Using https://en.wikipedia.org/wiki/Cramer%27s_rule
func solve(game Game) (int, int, bool) {

	/*
			buttonA.X * A + buttonB.X * B = prize.X
		    buttonA.Y * A + buttonB.Y * B = prize.Y

			a1 = buttonA.X , b1 = buttonB.X
			a2 = buttonA.Y , b2 = buttonB.Y

			c1 = prize.X, c2 = prize.Y


	*/
	buttonA := game.ButtonA
	buttonB := game.ButtonB
	prize := game.Prize

	// a1b2-b1a2
	denom := buttonA.X*buttonB.Y - buttonB.X*buttonA.Y

	// c1b2- b1c2
	xdiff := (prize.X*buttonB.Y - buttonB.X*prize.Y)

	// a1c2 - c1a2
	ydiff := (buttonA.X*prize.Y - prize.X*buttonA.Y)

	if denom != 0 && xdiff%denom == 0 && ydiff%denom == 0 {
		return xdiff / denom, ydiff / denom, true
	}

	return 0, 0, false
}

func getMinTokenCount(games []Game, include100Limit bool) int {
	sum := 0
	for _, game := range games {
		a, b, possible := solve(game)
		if possible && a > 0 && b > 0 {
			if include100Limit {
				if a <= 100 && b <= 100 {
					sum += (3*a + b)
				}
			} else {
				sum += (3*a + b)
			}
		}
	}
	return sum
}

func Map(games []Game, f func(g Game) Game) []Game {
	newGames := make([]Game, len(games))
	for _, game := range games {
		newGames = append(newGames, f(game))
	}
	return newGames
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	games := make([]Game, 0)

	values := make([]Point, 0)

	for scanner.Scan() {

		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "Prize") {
			prize := strings.Split(strings.Split(line, ":")[1], ",")

			x, y := strings.Split(prize[0], "=")[1], strings.Split(prize[1], "=")[1]

			xv, err := strconv.Atoi(x)
			if err != nil {
				log.Fatal(err)
			}

			yv, err := strconv.Atoi(y)
			if err != nil {
				log.Fatal(err)
			}

			games = append(games, Game{values[0], values[1], Point{xv, yv}})
			values = make([]Point, 0)

		} else {
			button := strings.Split(strings.Split(line, ":")[1], ",")

			ax, ay := strings.Split(button[0], "+")[1], strings.Split(button[1], "+")[1]

			x1, err := strconv.Atoi(ax)
			if err != nil {
				log.Fatal(err)
			}

			y1, err := strconv.Atoi(ay)
			if err != nil {
				log.Fatal(err)
			}
			values = append(values, Point{x1, y1})
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()

	fmt.Printf("Sum for part -1 : %d \n", getMinTokenCount(games, true))

	newGames := Map(games, func(g Game) Game {
		newPrize := Point{g.Prize.X + 10000000000000, g.Prize.Y + 10000000000000}
		return Game{g.ButtonA, g.ButtonB, newPrize}
	})

	fmt.Printf("Sum for part -2 : %d \n", getMinTokenCount(newGames, false))
}

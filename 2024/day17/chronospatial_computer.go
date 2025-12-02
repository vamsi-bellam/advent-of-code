package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func pow(a, b int64) int64 {

	result := int64(1)
	for b > 0 {
		result *= a
		b--
	}
	return result
}

type Memory struct {
	RegisterA, RegisterB, RegisterC int64
}

func (m *Memory) updateA(newVal int64) {
	m.RegisterA = newVal
}
func (m *Memory) updateB(newVal int64) {
	m.RegisterB = newVal
}
func (m *Memory) updateC(newVal int64) {
	m.RegisterC = newVal
}

func getComboOperand(op int64, mem Memory) int64 {
	if op >= 0 && op <= 3 {
		return op
	} else if op == 4 {
		return mem.RegisterA
	} else if op == 5 {
		return mem.RegisterB
	} else if op == 6 {
		return mem.RegisterC
	}
	return op
}
func runProgram(program []int64, memory Memory) []int64 {
	plen := len(program)
	i := 0
	output := make([]int64, 0)
	for i+1 < plen {
		ins := program[i]
		operand := program[i+1]
		if ins == 0 {
			res := memory.RegisterA / pow(2, getComboOperand(operand, memory))
			memory.updateA(res)
		} else if ins == 1 {
			memory.updateB(memory.RegisterB ^ operand)
		} else if ins == 2 {
			memory.updateB(getComboOperand(operand, memory) % 8)
		} else if ins == 3 {
			if memory.RegisterA != 0 {
				i = int(operand)
				continue
			}
		} else if ins == 4 {
			memory.updateB(memory.RegisterB ^ memory.RegisterC)
		} else if ins == 5 {
			output = append(output, getComboOperand(operand, memory)%8)
		} else if ins == 6 {
			res := memory.RegisterA / pow(2, getComboOperand(operand, memory))
			memory.updateB(res)
		} else if ins == 7 {
			res := memory.RegisterA / pow(2, getComboOperand(operand, memory))
			memory.updateC(res)
		} else {
			break
		}

		i = i + 2
	}
	return output
}

func join(values []int64, sep string) string {
	output := ""
	for i, ins := range values {
		output += strconv.Itoa(int(ins))
		if i != len(values)-1 {
			output += sep
		}
	}
	return output
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	memory := Memory{0, 0, 0}
	program := make([]int64, 0)
	for scanner.Scan() {

		line := scanner.Text()

		if strings.HasPrefix(line, "Register A") {
			num, err := strconv.Atoi(strings.Split(line, ": ")[1])
			if err != nil {
				log.Fatal(err)
			}
			memory.updateA(int64(num))
		} else if strings.HasPrefix(line, "Register B") {
			num, err := strconv.Atoi(strings.Split(line, ": ")[1])
			if err != nil {
				log.Fatal(err)
			}
			memory.updateB(int64(num))
		} else if strings.HasPrefix(line, "Register C") {
			num, err := strconv.Atoi(strings.Split(line, ": ")[1])
			if err != nil {
				log.Fatal(err)
			}
			memory.updateC(int64(num))
		} else if strings.HasPrefix(line, "Program") {
			nums := strings.Split(strings.Split(line, ": ")[1], ",")
			for _, num := range nums {
				val, err := strconv.Atoi(num)
				if err != nil {
					log.Fatal(err)
				}
				program = append(program, int64(val))
			}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()

	fmt.Printf("Output of the program is %s \n", join(program, ","))

	memory.updateA(1)

	// Found pattern from reddit community threads
	for {

		output := runProgram(program, memory)

		if reflect.DeepEqual(output, program) {
			break
		}

		if len(output) == len(program) {

			for j := len(program) - 1; j >= 0; j-- {
				if program[j] != output[j] {
					memory.updateA(memory.RegisterA + pow(8, int64(j)))
					break
				}
			}
		} else {
			memory.updateA(memory.RegisterA * 2)
		}
	}

	fmt.Printf("Lowest positive value of reg A to get same program is  %d \n", memory.RegisterA)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isSafeReport(report []int) bool {
	isIncreasing := report[1] > report[0]
	len := len(report)

	for i := 1; i < len; i++ {

		diff := report[i] - report[i-1]

		if (isIncreasing && (diff < 1 || diff > 3)) ||
			(!isIncreasing && (diff > -1 || diff < -3)) {
			return false
		}
	}

	return true
}

func isSafeReportWithSingleToleration(report []int) bool {
	len := len(report)

	if isSafeReport(report) {
		return true
	}

	for i := 0; i < len; i++ {
		newReport := make([]int, len)
		copy(newReport, report)

		newReport = append(newReport[:i], newReport[i+1:]...)

		if isSafeReport(newReport) {
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
	safeReports := 0
	safeReportsWithSingleToleration := 0

	for scanner.Scan() {
		line := scanner.Text()
		report := make([]int, 0)

		for _, v := range strings.Split(line, " ") {
			val, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			report = append(report, val)
		}

		if isSafeReport(report) {
			safeReports++
		}

		if isSafeReportWithSingleToleration(report) {
			safeReportsWithSingleToleration++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()
	fmt.Printf("Safe Reports - %d \n", safeReports)
	fmt.Printf("Safe Reports with single toleration - %d \n", safeReportsWithSingleToleration)
}

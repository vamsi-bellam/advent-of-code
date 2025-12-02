package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func checkSum(encodedData []int) int {
	checksum := 0

	for i, v := range encodedData {
		if v != -1 {
			checksum += i * v
		}
	}
	return checksum
}

func moveFiles(encodedData []int) {
	freeSpaceIndexFromLeft := 0
	encodedDataLen := len(encodedData)
	fileIndexFromRight := encodedDataLen - 1

	for i, v := range encodedData {
		if v == -1 {
			freeSpaceIndexFromLeft = i
			break
		}
	}

	for i := encodedDataLen - 1; i >= 0; i-- {
		if encodedData[i] != -1 {
			fileIndexFromRight = i
			break
		}
	}

	for fileIndexFromRight > freeSpaceIndexFromLeft {
		encodedData[freeSpaceIndexFromLeft] = encodedData[fileIndexFromRight]
		encodedData[fileIndexFromRight] = -1

		freeSpaceIndexFromLeft++
		fileIndexFromRight--

		for fileIndexFromRight > freeSpaceIndexFromLeft &&
			encodedData[freeSpaceIndexFromLeft] != -1 {
			freeSpaceIndexFromLeft++
		}

		for fileIndexFromRight > freeSpaceIndexFromLeft &&
			encodedData[fileIndexFromRight] == -1 {
			fileIndexFromRight--
		}
	}
}

func findNextFilePos(encodedData []int, start, end int) (int, int, bool) {

	fileEnd := -1
	fileStart := -1

	for i := end - 1; i >= start; i-- {
		if encodedData[i] != -1 {
			fileEnd = i
			break
		}
	}

	// if there is no end to file, then no start as well
	// => no file found in given range
	if fileEnd == -1 {
		return -1, -1, false
	}

	// if there is a file end then inititally file start is also same
	fileStart = fileEnd

	// then try to see if there is a start of file before end
	for i := fileEnd; i > start; i-- {
		if encodedData[i-1] == encodedData[i] {
			fileStart = i - 1
		} else {
			break
		}
	}
	return fileStart, fileEnd, true
}

func findNextFreeSpacePos(encodedData []int, start, end, fileLen int) (
	int, bool) {

	for start < end {
		freeSpaceStart := -1
		freeSpaceEnd := -1

		for i := start; i < end; i++ {
			if encodedData[i] == -1 {
				freeSpaceStart = i
				break
			}
		}

		// if there is no free space start then no free space end too
		if freeSpaceStart == -1 {
			return freeSpaceStart, false
		}

		// if there is a free space start then
		// inititally free space end is also same
		freeSpaceEnd = freeSpaceStart

		// then try to see if there is a end of free space after start
		for i := freeSpaceStart + 1; i < end; i++ {
			if encodedData[i] == encodedData[i-1] {
				freeSpaceEnd = i
			} else {
				break
			}
		}
		// if found free space have at least given fileLen then we found
		// required free space
		if freeSpaceEnd-freeSpaceStart+1 >= fileLen {
			return freeSpaceStart, true
		}

		// else check if there is a required free space in next blocks
		start = freeSpaceStart + 1
	}

	return -1, false
}

func moveFilesAtOnce(encodedData []int, fileCount int) {

	searchTill := len(encodedData)

	// for every file from end see if we can find enough free space to fit
	// entire file chunk and if possible move the file chunk to free space
	for fileCount > 0 {

		fileStart, fileEnd, foundFile :=
			findNextFilePos(encodedData, 0, searchTill)

		searchTill = fileStart

		// Ideally we will find file till the fileCount is over,
		// Just in case not found stop moving files
		if !foundFile {
			break
		}

		currFile := encodedData[fileStart]

		fileLen := fileEnd - fileStart + 1

		freeSpaceStart, foundSpace :=
			findNextFreeSpacePos(encodedData, 0, searchTill, fileLen)

		if foundSpace {
			for i := freeSpaceStart; i < freeSpaceStart+fileLen; i++ {
				encodedData[i] = currFile
			}

			for i := fileStart; i <= fileEnd; i++ {
				encodedData[i] = -1
			}
		}

		fileCount--
	}

}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	data := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		for _, v := range strings.Split(line, "") {
			num, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			data = append(data, num)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()

	id := 0

	encodedData := make([]int, 0)

	dataLen := len(data)

	for i := 0; i < dataLen; i++ {
		files := data[i]

		for files > 0 {
			encodedData = append(encodedData, id)
			files--
		}
		id++

		if i+1 < dataLen {
			freeSpace := data[i+1]
			for freeSpace > 0 {
				encodedData = append(encodedData, -1)
				freeSpace--
			}
			i++
		}
	}

	encodedDataCopy := make([]int, len(encodedData))

	copy(encodedDataCopy, encodedData)

	moveFiles(encodedData)
	fmt.Printf("File system checksum : %d \n", checkSum(encodedData))

	moveFilesAtOnce(encodedDataCopy, id)

	fmt.Printf("File system checksum2 : %d \n", checkSum(encodedDataCopy))
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func isStrInt(b byte) bool {
	return b < 58 && b > 48
}

func findNum(str string) int {
	size := len(str)
	chrToInt := map[string]int{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9}
	i, j := 0, size-1
	end1 := false
	end2 := false
	for i < size && j > 0 {
		if !isStrInt(str[i]) {
			i += 1
		} else {
			end1 = true
		}
		if !isStrInt(str[j]) {
			j -= 1
		} else {
			end2 = true
		}
		if end1 && end2 {
			break
		}
	}
	return chrToInt[string(str[i])]*10 + chrToInt[string(str[j])]
}

func findNum2(str string) int {
	mapToInt := map[string]int{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
		"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}
	keys := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9",
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	minIdx := math.MaxInt32
	maxIdx := math.MinInt32
	minKey := ""
	maxKey := ""
	for _, key := range keys {
		tmpMin := strings.Index(str, key)
		tmpMax := strings.LastIndex(str, key)
		if tmpMin < minIdx && tmpMin != -1 {
			minIdx = tmpMin
			minKey = key
		}
		if tmpMax > maxIdx && tmpMax != -1 {
			maxIdx = tmpMax
			maxKey = key
		}
	}
	return mapToInt[minKey]*10 + mapToInt[maxKey]
}

func main() {
	// open the file using Open() function from os library
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// read the file line by line using a scanner
	scanner := bufio.NewScanner(file)
	res1 := 0
	res2 := 0
	for scanner.Scan() {
		//Print the linej
		res1 += findNum(scanner.Text())
		res2 += findNum2(scanner.Text())
	}
	// check for the error that occurred during the scanning
	fmt.Println(res1)
	fmt.Println(res2)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the file at the end of the program
	file.Close()
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func isLegalGame(str string) (int, bool) {
	totals := []int{12, 13, 14}
	id, err := strconv.Atoi(strings.Split(strings.Split(str, ":")[0], " ")[1])
	ballType := map[string]int{"red": 0, "green": 1, "blue": 2}
	if err != nil {
		log.Fatal(err)
	}
	hands := strings.Split(strings.Split(str, ":")[1], ";")
	for _, v := range hands {
		balls := strings.Split(v, ",")
		for _, ball := range balls {
			tmp := strings.Split(ball, " ")
			vtmp, err := strconv.Atoi(tmp[1])
			if err != nil {
				log.Fatal(err)
			}
			if vtmp > totals[ballType[tmp[len(tmp)-1]]] {
				return id, false
			}
		}
	}
	return id, true
}

func powerOfGame(str string) (int, int) {
	maxColor := []int{math.MinInt, math.MinInt, math.MinInt}
	id, err := strconv.Atoi(strings.Split(strings.Split(str, ":")[0], " ")[1])
	ballType := map[string]int{"red": 0, "green": 1, "blue": 2}
	if err != nil {
		log.Fatal(err)
	}
	hands := strings.Split(strings.Split(str, ":")[1], ";")
	for _, v := range hands {
		balls := strings.Split(v, ",")
		for _, ball := range balls {
			tmp := strings.Split(ball, " ")
			vtmp, err := strconv.Atoi(tmp[1])
			if err != nil {
				log.Fatal(err)
			}
			if maxColor[ballType[tmp[len(tmp)-1]]] < vtmp {
				maxColor[ballType[tmp[len(tmp)-1]]] = vtmp
			}
		}
	}
	res := 1
	for _, v := range maxColor {
		res *= v
	}
	return id, res
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
		id, legal := isLegalGame(scanner.Text())
		id, pow := powerOfGame(scanner.Text())
		if legal {
			res1 += id
		}
		res2 += pow
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

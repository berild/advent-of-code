package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func findLegalNumbers(symb map[string]bool) int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// read the file line by line using a scanner
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile("[0-9]")
	neighs := []int{-1, 0, 1}
	res := 0
	var ok bool = false
	for row := 0; scanner.Scan(); row++ {
		//Print the linej
		tmp := scanner.Text()
		tmp2 := re.FindAllStringIndex(tmp, -1)
		tmpVal := re.FindAllString(tmp, -1)
		var tmpSeq []int
		var numb string
		for k, v := range tmp2 {
			if len(tmpSeq) == 0 {
				tmpSeq = append(tmpSeq, v[0])
				numb += tmpVal[k]
			} else if tmpSeq[len(tmpSeq)-1] == v[0]-1 {
				tmpSeq = append(tmpSeq, v[0])
				numb += tmpVal[k]
			} else {
				for _, i := range neighs {
					for j := tmpSeq[0] - 1; j < tmpSeq[len(tmpSeq)-1]+2; j++ {
						str := fmt.Sprintf("%d-%d", row+i, j)
						_, ok = symb[str]
						if ok {
							tmp3, err := strconv.Atoi(numb)
							if err != nil {
								log.Fatal(err)
							}
							res += tmp3
						}
						if i == 0 && j == tmpSeq[0]-1 {
							j = tmpSeq[len(tmpSeq)-1]
						}
					}
				}
				numb = tmpVal[k]
				tmpSeq = []int{v[0]}
			}
			if k == len(tmp2)-1 {
				for _, i := range neighs {
					for j := tmpSeq[0] - 1; j < tmpSeq[len(tmpSeq)-1]+2; j++ {
						str := fmt.Sprintf("%d-%d", row+i, j)
						_, ok = symb[str]
						if ok {
							tmp3, err := strconv.Atoi(numb)
							if err != nil {
								log.Fatal(err)
							}
							res += tmp3
						}
						if i == 0 && j == tmpSeq[0]-1 {
							j = tmpSeq[len(tmpSeq)-1]
						}
					}
				}
			}
		}
	}
	// check for the error that occurred during the scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the file at the end of the program
	file.Close()
	return res
}

func finGearRatio(symb map[string][]int) int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// read the file line by line using a scanner
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile("[0-9]")
	neighs := []int{-1, 0, 1}
	var ok bool = false
	for row := 0; scanner.Scan(); row++ {
		//Print the linej
		tmp := scanner.Text()
		tmp2 := re.FindAllStringIndex(tmp, -1)
		tmpVal := re.FindAllString(tmp, -1)
		var tmpSeq []int
		var numb string
		for k, v := range tmp2 {
			if len(tmpSeq) == 0 {
				tmpSeq = append(tmpSeq, v[0])
				numb += tmpVal[k]
			} else if tmpSeq[len(tmpSeq)-1] == v[0]-1 {
				tmpSeq = append(tmpSeq, v[0])
				numb += tmpVal[k]
			} else {
				for _, i := range neighs {
					for j := tmpSeq[0] - 1; j < tmpSeq[len(tmpSeq)-1]+2; j++ {
						str := fmt.Sprintf("%d-%d", row+i, j)
						_, ok = symb[str]
						if ok {
							tmp3, err := strconv.Atoi(numb)
							if err != nil {
								log.Fatal(err)
							}
							symb[str][0] += 1
							symb[str][1] *= tmp3
						}
						if i == 0 && j == tmpSeq[0]-1 {
							j = tmpSeq[len(tmpSeq)-1]
						}
					}
				}
				numb = tmpVal[k]
				tmpSeq = []int{v[0]}
			}
			if k == len(tmp2)-1 {
				for _, i := range neighs {
					for j := tmpSeq[0] - 1; j < tmpSeq[len(tmpSeq)-1]+2; j++ {
						str := fmt.Sprintf("%d-%d", row+i, j)
						_, ok = symb[str]
						if ok {
							tmp3, err := strconv.Atoi(numb)
							if err != nil {
								log.Fatal(err)
							}
							symb[str][0] += 1
							symb[str][1] *= tmp3
						}
						if i == 0 && j == tmpSeq[0]-1 {
							j = tmpSeq[len(tmpSeq)-1]
						}
					}
				}
			}
		}
	}
	// check for the error that occurred during the scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the file at the end of the program
	file.Close()
	res := 0
	for _, val := range symb {
		if val[0] == 2 {
			res += val[1]
		}
	}
	return res
}

func findSymbols() (map[string]bool, map[string][]int) {
	symbols := make(map[string]bool)
	symbols2 := make(map[string][]int)
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// read the file line by line using a scanner
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile("[^0-9.]")
	// re2 := regexp.MustCompile("[*]")

	for row := 0; scanner.Scan(); row++ {
		//Print the linej
		tmp := scanner.Text()
		tmp2 := re.FindAllStringIndex(tmp, -1)
		for _, v := range tmp2 {
			str := fmt.Sprintf("%d-%d", row, v[0])
			symbols[str] = true
			if tmp[v[0]:v[1]] == "*" {
				symbols2[str] = []int{0, 1}
			}
		}
	}
	// check for the error that occurred during the scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the file at the end of the program
	file.Close()
	return symbols, symbols2
}

func main() {
	symbols, symbols2 := findSymbols()
	res := findLegalNumbers(symbols)
	res2 := finGearRatio(symbols2)
	fmt.Println(res)
	fmt.Println(res2)
}

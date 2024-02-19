package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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
		scanner.Text()
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

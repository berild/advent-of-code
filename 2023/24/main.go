package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Hail struct {
	a    float64
	b    float64
	c    float64
	sx   float64
	sy   float64
	sz   float64
	vx   float64
	vy   float64
	vz   float64
	next *Hail
}

type Hails struct {
	head *Hail
}

func str2int(arr []string) []float64 {
	var res []float64
	for _, v := range arr {
		tmp, err := strconv.Atoi(strings.ReplaceAll(v, " ", ""))
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, float64(tmp))
	}
	return res
}

func cHail(str string) *Hail {
	tmp := str2int(strings.Split(strings.ReplaceAll(str, "@", ","), ","))
	c := tmp[0]*tmp[4] - tmp[1]*tmp[3]
	return &Hail{a: tmp[4], b: -tmp[3], c: c, sx: tmp[0], sy: tmp[1], sz: tmp[2], vx: tmp[3], vy: tmp[4], vz: tmp[5]}
}

func testHail(hs1 *Hail, hs2 *Hail) int {
	a1, b1, c1 := hs1.a, hs1.b, hs1.c
	a2, b2, c2 := hs2.a, hs2.b, hs2.c
	if a1*b2 == b1*a2 {
		return 0
	}
	x := (c1*b2 - c2*b1) / (a1*b2 - a2*b1)
	y := (c2*a1 - c1*a2) / (a1*b2 - a2*b1)
	if 200000000000000 <= x && x <= 400000000000000 && 200000000000000 <= y && y <= 400000000000000 {
		if (x-hs1.sx)*hs1.vx >= 0 && (y-hs1.sy)*hs1.vy >= 0 && (x-hs2.sx)*hs2.vx >= 0 && (y-hs2.sy)*hs2.vy >= 0 {
			return 1
		}
	}
	return 0
}

func (list *Hails) Insert(str string) int {
	newHail := cHail(str)
	count := 0
	if list.head == nil {
		list.head = newHail
	} else {
		current := list.head
		for current.next != nil {
			count += testHail(current, newHail)
			current = current.next
		}
		count += testHail(current, newHail)
		current.next = newHail
	}
	return count
}

func main() {
	file, err := os.Open("input.txt")
	count := 0
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	list := Hails{}
	for scanner.Scan() {
		count += list.Insert(scanner.Text())
	}
	fmt.Println(count)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()
}

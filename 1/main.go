package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	var x = 0
	var dup = false
	var one = true
	var h = make(map[int]int)
	history(h, x)
	for !dup {
		x, dup = read(h, x)
		if one {
			fmt.Printf("pt1 output is %d \n", x)
			one = false
		}
	}
	fmt.Printf("pt2 output is %d \n", x)

}
func history(m map[int]int, current int) bool {
	_, ok := m[current]
	if ok {
		return true
	} else {
		m[current] = 1
		return false
	}
}
func read(m map[int]int, current int) (int, bool) {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var in = scanner.Text()
		var y, err = strconv.Atoi(in)
		if err != nil {
			log.Fatal(err)
		}
		current = current + y
		if history(m, current) {
			return current, true
		}
	}
	return current, false
}

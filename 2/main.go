package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var z [2]int
	var set [] string
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var in = scanner.Text()
		set=append(set, in)
		two, three := check(in)
		if two {
			z[0] = z[0] + 1
		}
		if three {
			z[1] = z[1] + 1
		}
	}
	var nearmiss =""
	for x:=0;x< len(set);x++{
		for y:=x+1;y<len(set);y++ {
			wmatch :=0
			w:= ""
			fmt.Printf("%v \n", set[y] )
			for s:=0;s<len(set[x]);s++{

				if ( string([]rune(set[x])[s]) == string([]rune(set[y])[s])  ){
					wmatch++
					w = w + string([]rune(set[x])[s])
				}
				if wmatch == len(set[x])-1{
					nearmiss = w
				}
			}

		}
	}

	fmt.Printf(" two: %d \n three: %d \n pt1 hash: %d \n",z[0],z[1],z[0]*z[1])
	fmt.Printf("part two %v",nearmiss)


}

/*
To make sure you didn't miss any, you scan the likely candidate boxes again, counting the number that have an ID containing exactly
two of any letter and then separately counting those with exactly three of any letter. You can multiply those two counts together to
get a rudimentary checksum and compare it to what your device predicts.
 */



func check(in string) (bool,bool) {
	var two = false
	var three = false
	var h = make(map[string]int)
	for x:=0;x<len(in) ;x ++  {
		c := string([]rune(in)[x])
		y,ok := h[c]
		if !ok{
			y = 0
		}
		h[c]= y+1
	}
	for _, v := range h {
		switch v {
		case 2:
			two = true
		case 3:
			three = true
		}
	}


	return two,three


}
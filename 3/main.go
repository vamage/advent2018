package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type patch struct {
	Item   int
	Pointx int
	Pointy int
	Sizex  int
	Sizey  int
}

func main() {
	var patches []patch
	grid := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		grid[i] = make([]int, 1000)
	}
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var in = scanner.Text()
		var item, pointx, pointy, sizex, sizey int
		fmt.Sscanf(in, "#%d @ %d,%d: %dx%d", &item, &pointx, &pointy, &sizex, &sizey)
		fmt.Printf("#%v  %v,%v  b %v,%v \n", item, pointx, pointy, sizex, sizey)
		patches = append(patches, patch{item, pointx, pointy, sizex, sizey})
		update(grid, pointx, pointy, sizex, sizey)

	}
	o := check(grid)
	olap := overlap(grid, patches)
	fmt.Printf("pt1 overlap is %v \n", o)
	fmt.Printf("pt2 nonoverlap is %v \n",olap)

}

func update(grid [][]int, pointx, pointy, sizex, sizey int) {

	for x := 0; x < sizex; x++ {
		for y := 0; y < sizey; y++ {
			grid[pointx+x][pointy+y] = grid[pointx+x][pointy+y] + 1
		}
	}

}

func check(grid [][]int) int {
	overlap := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if grid[x][y] > 1 {
				overlap++
			}
		}

	}
	return overlap

}

func overlap(grid [][]int, patches [] patch) int {
	for p := 0; p < len(patches); p++ {
		patch := patches[p]
		overlaping:=false
		for x := 0; x < patch.Sizex; x++ {
			for y := 0; y < patch.Sizey; y++ {
				if  grid[patch.Pointx+x][patch.Pointy+y] >1 {
					overlaping = true
				}
			}
		}
		if !overlaping {
			return patch.Item

		}
	}
	return -1
}

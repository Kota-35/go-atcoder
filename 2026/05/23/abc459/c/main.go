// https://atcoder.jp/contests/abc459/tasks/abc459_c
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isPiledAllField(grid []int, n int) bool { 
	for i := range n {
		if grid[i] == 0 {
			return false
		}
	}

	return true
}

func processX(grid []int, x int, n int) {
	grid[x-1] += 1

	if isPiledAllField(grid, n){
		for i := range n {
			grid[i]-= 1
		}	
	}
}


func processY(grid []int, y int, n int) {
	count := 0
	for i := range n {
		if grid[i] >= y {
			count++
		}	
	} 
	fmt.Println(count)
}

func main() {
	var n, q int
	fmt.Scan(&n, &q)
	sc := bufio.NewScanner(os.Stdin)
	grid := make([]int, n)
	for range q{
		sc.Scan()
		inputs := strings.Split(sc.Text(), " ")
		first, _ := strconv.Atoi(inputs[0])
		seconde, _ := strconv.Atoi(inputs[1])
		if first == 1 {
			processX(grid, seconde, n)
		} else if first == 2 {
			processY(grid, seconde, n)
		}
	}



}

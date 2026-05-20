// https://atcoder.jp/contests/abc318/tasks/abc318_b
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	grid := make([][]bool, 100)
	for i := range 100 {
		grid[i] = make([]bool, 100)
	}
	var n int
	fmt.Scan(&n)

	sc := bufio.NewScanner(os.Stdin)
	for range n {
		sc.Scan()
		inputs := strings.Split(sc.Text(), " ")
		a, _ := strconv.Atoi(inputs[0])
		b, _ := strconv.Atoi(inputs[1])
		c, _ := strconv.Atoi(inputs[2])
		d, _ := strconv.Atoi(inputs[3])

		for i := a; i < b; i++ {
			for j := c; j < d; j++ {
				grid[i][j] = true
			}
		}
	}

	ans := 0
	for i := range 100 {
		for j := range 100 {
			if grid[i][j] {
				ans += 1
			}
		}
	}

	fmt.Println(ans)

}

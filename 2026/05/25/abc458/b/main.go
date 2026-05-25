// https://atcoder.jp/contests/abc458/tasks/abc458_b
package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	matrix := make([][]int, h)
	for i := range h {
		matrix[i] = make([]int, w)
	}

	for i := range h {
		for j := range w {
			if (i-1 >= 0) {
				matrix[i][j]++
			}

			if (i+1 < h) {
				matrix[i][j]++
			}

			if (j-1 >= 0) {
				matrix[i][j]++
			}

			if (j+1 < w) {
				matrix[i][j]++
			}
		}
	} 

	for i := range h {
		for j := range w {
			if j == w-1 {
				fmt.Printf("%d\n", matrix[i][j])
			} else {
				fmt.Printf("%d ", matrix[i][j])
			}
		}
	}
}

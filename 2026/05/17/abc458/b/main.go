// https://atcoder.jp/contests/abc458/tasks/abc458_b
package main

import (
	"fmt"
)

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	matrix := make([][]int, h)
	for i := 0; i < h; i++ {
		matrix[i] = make([]int, w)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if i-1 >= 0 {
				matrix[i][j] += 1
			}

			if i+1 < h {
				matrix[i][j] += 1
			}

			if j-1 >= 0 {
				matrix[i][j] += 1
			}

			if j+1 < w {
				matrix[i][j] += 1
			}
		}
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if j != w-1 {
				fmt.Printf("%d ", matrix[i][j])
			} else {
				fmt.Printf("%d", matrix[i][j])
			}
		}
		fmt.Println()
	}
}

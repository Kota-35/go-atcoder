// https://atcoder.jp/contests/abc418/tasks/abc418_b
package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)

	ans := 0.0
	n := len(s)

	for i := range n {
		if s[i] != 't' {
			continue
		}

		for j := i + 2; j < n; j++ {
			if s[j] != 't' {
				continue
			}

			x := 0
			for ki := i; ki <= j; ki++ {
				if s[ki] == 't' {
					x++
				}
			}
			length := j - i + 1

			item := float64(x-2) / float64(length-2)
			if item > ans {
				ans = item
			}
		}
	}

	fmt.Printf("%.17f\n", ans)
}

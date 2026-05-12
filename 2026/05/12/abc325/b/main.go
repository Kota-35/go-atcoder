// https://atcoder.jp/contests/abc325/tasks/abc325_b
package main

import (
	"bufio"
	"fmt"
	"os"
)

// 各拠点 i が参加できる UTC 開始時刻 T の条件: 9 - X_i <= T <= 17 - X_i
// X_i は 0~23 の整数なので T の範囲は -14 ~ 17

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	read := func() int {
		sc.Scan()
		var v int
		fmt.Sscan(sc.Text(), &v)
		return v
	}

	n := read()
	ws := make([]int, n)
	xs := make([]int, n)
	for i := 0; i < n; i++ {
		ws[i] = read()
		xs[i] = read()
	}

	ans := 0
	for t := -14; t <= 17; t++ {
		sum := 0
		for i := 0; i < n; i++ {
			local := t + xs[i]
			if local >= 9 && local <= 17 {
				sum += ws[i]
			}
		}
		if sum > ans {
			ans = sum
		}
	}

	fmt.Println(ans)
}

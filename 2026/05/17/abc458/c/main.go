// https://atcoder.jp/contests/abc458/tasks/abc458_c
package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	count := 0
	for i := range len(s) {
		if s[i] == 'C' {
			// min(i, len(s)-i-1) + 1 の最後の`+1` は
			// 部分文字列が`C`の時の分のカウント
			count += min(i, len(s)-i-1) + 1
		}
	}

	fmt.Println(count)
}

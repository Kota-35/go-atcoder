// https://atcoder.jp/contests/abc459/tasks/abc459_d
package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader
var writer *bufio.Writer

func solve(s string) {
	// 各文字の出現回数を数える
	d := [26]int{}
	for _, c := range s {
		d[c-'a']++
	}

	// 最貧文字が過半数を超えるか判定
	n := len(s)
	for _, cnt := range d {
		if cnt*2 > n+1 {
			fmt.Fprintln(writer, "No")
			return
		}
	}

	fmt.Fprintln(writer, "Yes")

	// 構築
	result := make([]byte, 0, n)
	last := -1 // 直前に使った文字のインデックス
	
	for range n {
		// last 以外でd[c]が最大のcを選ぶ
		best := -1
		for c := range 26 {
			if c == last {
				continue
			}
			if d[c]== 0 {
				continue
			}
			if best== -1 || d[c] > d[best] {
				best =c 
			}
		}
		result = append(result, byte('a'+best))
		d[best]--
		last = best
	}

	fmt.Fprintf(writer, "%s\n", result)
}

func main() {
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)

	for t > 0{
		var s string 
		fmt.Fscan(reader, &s)
		solve(s)
		t--
	}
}

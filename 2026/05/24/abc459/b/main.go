// https://atcoder.jp/contests/abc459/tasks/abc459_b
package main

import (
	"fmt"
	"strings"
)

func main() {
	// アルファベット26文字からそれぞれに対応するCの数を並べたもの
	StoC := "22233344455566677778889999"
	var n int 
	fmt.Scan(&n)
	// += は非効率
	// Goの文字列はイミュータブルなので連結のたびに新しいメモリ確保+コピーが走るため
	var sb strings.Builder
	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		// s[0]-'a'でインんデックスを計算
		sb.WriteByte(StoC[s[0]-'a'])
	}
	fmt.Println(sb.String())
}

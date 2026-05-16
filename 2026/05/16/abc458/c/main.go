// https://atcoder.jp/contests/abc458/tasks/abc458_c
package main

import "fmt"

func main() {
	// コンテスト時は全ての部分文字列を全探索し、対象の文字列かどうかを判定するロジックで
	// O(n^2)と遅い
	// var s string
	// fmt.Scan(&s)

	// count := 0

	// for left := 0; left < len(s); left++ {
	// 	for right := left + 1; right < len(s)+1; right++ {
	// 		tmp := s[left:right]

	// 		l := len(tmp)
	// 		half := (l + 1) / 2

	// 		if l%2 == 1 && tmp[half-1] == 'C' {
	// 			count++
	// 		}
	// 	}
	// }

	// fmt.Println(count)

	// 解答
	// S のうち i 文字目に`C`があるとき、その`C`が真ん中に来る部分文字列を数える
	// `C`の左側には i-1 個, 右側には |S|-i 個の文字がある
	// その結果、`C`を真ん中とする文字列として、`C`の左右それぞれに0,1...,min(i-1, |S|-i)個文字を取り付けた
	// 合計 min(i, |S|-i+1)個の文字列を得ることができる.
	// これを、S中に存在する文字`C`全てについて足し合わせれば良い

	var s string
	fmt.Scan(&s)
	count := 0

	l := len(s)
	for i := 0; i < l; i++ {
		if s[i] == 'C' {
			// 例えば、
			//
			// S = "XAXCBXB"
			//
			// を考える
			// 右側: i = 3 文字 (X, A, X)
			// 左側: l - i - 1 = 3文字 (B, X, B)
			//
			// `C`が中心になるには, 左右に同じ数k個の文字を取る必要がある.
			//
			// k = 0: "C"        <- 1通り
			// k = 1: "XCB"      <- 1通り
			// k = 2: "AXCBX"    <- 1通り
			// k = 3: "XAXCBXB"  <- 1通り
			//
			// 取れる k の値の範囲は, 0, 1, ..., min(i, l-i-1)
			count += min(i, l-i-1) + 1
		}

	}
	fmt.Println(count)
}

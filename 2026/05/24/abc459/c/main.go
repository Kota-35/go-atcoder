// https://atcoder.jp/contests/abc459/tasks/abc459_c
//
// 解法のポイント:
// 「全マスから1個ずつ削除する」操作を実際には行わず、代わりに
// 全マスのブロック数の最小値 mn (= B[j]=N となる最大のj) を管理する。
//
// クエリ2で y 個以上のブロックが積まれたマス数を求めるとき、
// 実際の最小値が mn なので、答えは「(y+mn) 個以上のブロックがあるマス数」= c[y+mn] に等しい。
//
// mn はクエリ1のたびに高々1しか増えないため、O(Q) で管理できる。
// 全体計算量: O(N+Q)
package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAXN = 300001

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, q int
	fmt.Fscan(reader, &n, &q)

	a := make([]int, MAXN) // a[x]: マスxのブロック数(削除操作なしの累積値)
	// c[j]: j個以上ブロックが積まれているマスの数
	// a[x]は単調増加(一度追加したブロックは消えない)するため、c[j]++した後にc[j]--される場面が存在しない
	c := make([]int, MAXN)
	mn := 0                // 全マスのブロック数の最小値 = c[j]=N となる最大のj

	for range q {
		var t, x int
		fmt.Fscan(reader, &t, &x)
		if t == 1 {
			a[x]++
			c[a[x]]++
			// c[a[x]]==n: 全マスがa[x]個以上になった → 最小値がa[x]に更新
			if c[a[x]] == n {
				mn = a[x]
			}
		} else {
			// 削除しない世界での「y+mn 個以上のマス数」が答え
			if x+mn > q {
				fmt.Fprintln(writer, 0)
			} else {
				fmt.Fprintln(writer, c[x+mn])
			}
		}
	}
}

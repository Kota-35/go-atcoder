// https://atcoder.jp/contests/abc325/tasks/abc325_c
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	sc := bufio.NewScanner(os.Stdin)
	// bufio.Scanner のデフォルトバッファは 64KB. この問題の制約は W <= 1000 なので、１行最大1000文字であり余裕で収まる
	// １行が 64KB を超えるケース, 例えば、 W = 100000 のような問題の時に以下が必要
	//	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	ss := make([]string, h)
	for i := range h {
		sc.Scan()
		ss[i] = sc.Text()
	}

	visited := make([][]bool, h)
	for i := range h {
		visited[i] = make([]bool, w)
	}

	// ８方向の移動ベクトルを表す配列
	// {-1, -1} {-1, 0} {-1, 1}
	// { 0, -1}					{ 0, 1}
	// { 1, -1} { 1, 0} { 1, 1}
	//
	// 現在地(x, y)に対して、dirsの各要素(dx, dy)を足すことで
	// 隣接8マスの座標(x+dx, y+dy)が得られる
	dirs := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	ans := 0
	for i := range h {
		for j := range w {
			if ss[i][j] == '.' || visited[i][j] {
				continue
			}

			// '#'かつ未訪問のマスを見つけた
			// -> まだどのグループにも属していない新しいセンサを発見
			visited[i][j] = true
			queue := [][2]int{{i, j}}

			for len(queue) > 0 {
				curr := queue[0]
				queue = queue[1:]

				// 八近傍について調べる.
				for _, d := range dirs {
					nx := curr[0] + d[0]
					ny := curr[1] + d[1]
					// 隣接する座標がマスの外
					if nx < 0 || nx >= h || ny < 0 || ny >= w {
						continue
					}
					// マスの近傍のセンサが配置されている かつ その近傍はまだ未探索
					if ss[nx][ny] == '#' && !visited[nx][ny] {
						visited[nx][ny] = true
						// queueにその近傍を追加
						queue = append(queue, [2]int{nx, ny})
					}
				}
			}

			// 連結成分を１つ発見したのでインクリメント
			ans++
		}
	}

	fmt.Println(ans)
}

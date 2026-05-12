// https://atcoder.jp/contests/abc325/tasks/abc325_c
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var h, w int
	fmt.Fscan(reader, &h, &w)

	ss := make([]string, h)
	for i := 0; i < h; i++ {
		fmt.Fscan(reader, &ss[i])
	}

	visited := make([][]bool, h)
	for i := range visited {
		visited[i] = make([]bool, w)
	}

	dirs := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	bfs := func(si, sj int) {
		queue := [][2]int{{si, sj}}
		visited[si][sj] = true
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			for _, d := range dirs {
				ni, nj := cur[0]+d[0], cur[1]+d[1]
				if ni < 0 || ni >= h || nj < 0 || nj >= w {
					continue
				}
				if visited[ni][nj] || ss[ni][nj] != '#' {
					continue
				}
				visited[ni][nj] = true
				queue = append(queue, [2]int{ni, nj})
			}
		}
	}

	count := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if ss[i][j] == '#' && !visited[i][j] {
				bfs(i, j)
				count++
			}
		}
	}

	fmt.Println(count)
}

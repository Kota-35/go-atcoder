package main

import "fmt"

func bfs(graph map[string][]string, startNode string) {
	// 訪問済みのノードを追跡するマップ
	visited := make(map[string]bool)

	// queue としてスライスを使用
	queue := []string{startNode}
	visited[startNode] = true

	// キューが空になるまでループ
	for len(queue) > 0 {
		// 先頭のノードを乗り出す
		currentNode := queue[0]
		queue = queue[1:]

		fmt.Printf("訪問: %s\n", currentNode)

		// 隣接するノードをチェック
		for _, neighbor := range graph[currentNode] {
			if !visited[neighbor] {
				visited[neighbor] = true
				// キューの末尾に追加
				queue = append(queue, neighbor)
			}
		}
	}
}

func main() {
	graph := map[string][]string{
		"A": {"B", "C"},
		"B": {"A", "D", "E"},
		"C": {"A", "F"},
		"D": {"B"},
		"E": {"B", "F"},
		"F": {"C", "E"},
	}

	bfs(graph, "A")
}

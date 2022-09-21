// DFS와 BFS: https://www.acmicpc.net/problem/1260
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	out     = bufio.NewWriter(os.Stdout)
	in      = bufio.NewReader(os.Stdin)
	conn    = map[int][]int{}
	visited = map[int]bool{}
)

func main() {
	// readFile()
	defer out.Flush()
	var N, M, V int
	fmt.Fscanf(in, "%d %d %d\n", &N, &M, &V)
	for i := 0; i < M; i++ {
		var a, b int
		fmt.Fscanf(in, "%d %d\n", &a, &b)
		conn[a] = append(conn[a], b)
		conn[b] = append(conn[b], a)
	}

	for _, arr := range conn {
		sort.Ints(arr)
	}
	dfs(V)
	fmt.Fprintln(out)
	visited = map[int]bool{}
	bfs(N, V)
}

func dfs(here int) {
	if visited[here] {
		return
	}
	visited[here] = true
	fmt.Fprintf(out, "%d ", here)
	for _, v := range conn[here] {
		dfs(v)
	}
}

func bfs(nodeNum, here int) {
	result := make([]int, 0, nodeNum)
	que := make([]int, 0, nodeNum)
	que = append(que, here)
	for {
		if len(que) == 0 {
			break
		}
		here = que[0]
		que = que[1:]
		if !visited[here] {
			result = append(result, here)
		}
		for _, v := range conn[here] {
			if !visited[v] {
				que = append(que, v)
			}
			visited[here] = true
		}
	}
	for _, v := range result {
		fmt.Fprintf(out, "%d ", v)
	}
}

// func readFile() {
// 	f, err := os.Open("./input01.txt")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	r = bufio.NewReader(f)
// }

// 네트워크: https://school.programmers.co.kr/learn/courses/30/lessons/43162
// DFS: https://velog.io/@falling_star3/2.-깊이우선탐색DFS과-넓이우선탐색BFS
// python: https://velog.io/@uoayop/Programmers-43162-네트워크Python
package main

import "fmt"

var visited map[int]bool

func solution(n int, computers [][]int) int {
	group := 0
	visited = map[int]bool{}
	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs(i, computers)
			group += 1
		}
	}
	return group
}

func dfs(idx int, computers [][]int) {
	if !visited[idx] {
		visited[idx] = true
		// for j := idx; j < len(computers); j++ { 이걸로 될 거 같은데 왜 안될까....
		for j := 0; j < len(computers); j++ {
			if computers[idx][j] == 1 {
				computers[idx][j] = 0
				dfs(j, computers)
			}
		}
	}
}

func main() {
	ans1 := solution(3, [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}})
	fmt.Println(">>", 2, ans1)

	ans2 := solution(3, [][]int{{1, 1, 0}, {1, 1, 1}, {0, 1, 1}})
	fmt.Println(">>", 1, ans2)

	ans3 := solution(5, [][]int{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 1, 0, 0}, {0, 0, 0, 1, 0}, {0, 0, 0, 0, 1}})
	fmt.Println(">>", 4, ans3)
}

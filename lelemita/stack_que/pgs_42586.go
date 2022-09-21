//https://school.programmers.co.kr/learn/courses/30/lessons/42586

package main

import (
	"fmt"
)

func main() {
	fmt.Println(solution([]int{93, 30, 55}, []int{1, 30, 5}))                     // [2, 1]
	fmt.Println(solution([]int{95, 90, 99, 99, 80, 99}, []int{1, 1, 1, 1, 1, 1})) // [1, 3, 2]
}

func solution(progresses []int, speeds []int) []int {
	result := []int{}
Outer:
	for len(progresses) > 0 {
		turn := (100 - progresses[0]) / speeds[0]
		under := (100 - progresses[0]) % speeds[0]
		if under != 0 {
			turn += 1
		}
		for j := 0; j < len(progresses); j++ {
			progresses[j] += turn * speeds[j]
		}
		for j := 0; j < len(progresses); j++ {
			if progresses[j] < 100 {
				result = append(result, j)
				progresses = progresses[j:]
				speeds = speeds[j:]
				break
			}
			if j == len(progresses)-1 {
				result = append(result, len(progresses))
				break Outer
			}
		}
	}
	return result
}

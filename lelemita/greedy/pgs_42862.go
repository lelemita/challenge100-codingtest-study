// https://school.programmers.co.kr/learn/courses/30/lessons/42862
package main

import (
	"fmt"
	"sort"
)

func solution(n int, lost []int, reserve []int) int {
	absent := 0
	sort.Ints(lost)
	sort.Ints(reserve)

	newLost := []int{}
	for _, l := range lost {
		if rdx, exist := findIdx(l, reserve); exist {
			reserve = append(reserve[:rdx], reserve[rdx+1:]...)
		} else {
			newLost = append(newLost, l)
		}
	}
	lost = newLost

	for _, l := range lost {
		rental := false
		for _, num := range []int{0, -1, 1} {
			if idx, exist := findIdx(l+num, reserve); exist {
				reserve = append(reserve[:idx], reserve[idx+1:]...)
				rental = true
				break
			}
		}
		if !rental {
			absent += 1
		}
	}
	return n - absent
}

func findIdx(n int, arr []int) (int, bool) {
	for idx, v := range arr {
		if v == n {
			return idx, true
		}
	}
	return -1, false
}

func main() {
	fmt.Println(solution(5, []int{2, 4}, []int{1, 3, 5}), 5)
	fmt.Println(solution(5, []int{2, 4}, []int{3}), 4)
	fmt.Println(solution(3, []int{3}, []int{1}), 2)
	fmt.Println(solution(3, []int{1, 2, 3}, []int{1, 2, 3}), 3)

	// hint_case_1
	fmt.Println(solution(4, []int{2, 3}, []int{3, 4}), 3)
	// hint_case_2
	fmt.Println(solution(5, []int{4, 2}, []int{3, 5}), 5)
	// hint_case_3
	fmt.Println(solution(6, []int{3, 4, 5}, []int{3, 4, 6}), 6)

	// run time error
	fmt.Println(solution(6, []int{1, 6}, []int{1, 6}), 4)
}

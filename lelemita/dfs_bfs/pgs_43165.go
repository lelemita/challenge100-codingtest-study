package main

import "fmt"

var result int = 0
var target int
var numbers []int

func solution(nums []int, t int) int {
	result = 0
	target = t
	numbers = nums
	dfs(0, 0)
	return result
}

func dfs(idx, res int) {
	if idx == len(numbers) {
		if res == target {
			result += 1
		}
		return
	}

	dfs(idx+1, res+numbers[idx])
	dfs(idx+1, res-numbers[idx])
}

func main() {
	fmt.Println(solution([]int{1, 1, 1, 1, 1}, 3)) // 5
	fmt.Println(solution([]int{4, 1, 2, 1}, 4))    // 2
}

/* better solution

func solution(nums []int, target int) int {
	if len(nums) == 0 {
		if target == 0 {
			return 1
		} else {
			return 0
		}
	}
	return solution(nums[1:], target+nums[0]) + solution(nums[1:], target-nums[0])
}
*/

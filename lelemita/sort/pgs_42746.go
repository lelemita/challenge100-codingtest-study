// pgs_42746
package main

import (
	"fmt"
	"sort"
	"strconv"
)

// []string의 Alias
type Strs []string

func (s Strs) Len() int {
	return len(s)
}

func (s Strs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Strs) Less(i, j int) bool {
	return s[i]+s[j] > s[j]+s[i]
}

func solution(numbers []int) string {
	strs := []string{}
	isAllZero := true
	for _, n := range numbers {
		if isAllZero && n != 0 {
			isAllZero = false
		}
		strs = append(strs, strconv.Itoa(n))
	}
	if isAllZero {
		return "0"
	}
	sort.Sort(Strs(strs)) // casting

	result := ""
	for _, v := range strs {
		result += v
	}
	return result
}

func main() {
	fmt.Println(solution([]int{6, 10, 2}) == "6210")
	fmt.Println(solution([]int{3, 30, 34, 5, 9}) == "9534330")
}

/* better solution

import "sort"
import "strconv"

func solution(numbers []int) string {
    sort.Slice(numbers, func(i, j int) bool {
        s1 := strconv.Itoa(numbers[i])
        s2 := strconv.Itoa(numbers[j])
        return s1+s2 > s2+s1
    })

    if numbers[0] == 0 {
        return "0"
    }
    answer := ""
    for i := range numbers {
        answer += strconv.Itoa(numbers[i])
    }
    return answer
}

*/

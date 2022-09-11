package main

import (
	"fmt"
)

func solution(s string) string {
	result := []rune{}
	isWordStart := true
	for _, letter := range s {
		if isWordStart {
			if letter <= 'z' && letter >= 'a' {
				letter += ('A' - 'a')
			}
		} else {
			if letter <= 'Z' && letter >= 'A' {
				letter -= ('A' - 'a')
			}
		}
		isWordStart = (letter == ' ')
		result = append(result, letter)
	}
	return string(result)
}

func main() {
	fmt.Println(solution("3people   unFollowed me"))
	fmt.Println(solution("for the last week"))
}

// 자괴감 들었던 다른 사람의 풀이...ㅠㅠ
// import "strings"

// func solution(s string) string {
//     return strings.Title(strings.ToLower(s))
// }

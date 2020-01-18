package main

import "fmt"

func scoreOfParentheses(S string) int {
	var current int
	stack := make([]int, 0)
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			stack = append(stack, current)
			current = 0
		} else {
			current = stack[len(stack)-1] + max(current*2, 1)
			stack = stack[:len(stack)-1]
		}
	}
	return current
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "(()(()))"
	fmt.Println(scoreOfParentheses(s))
}

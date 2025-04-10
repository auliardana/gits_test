package main

import (
	"fmt"
)

func isBalanced(s string) string {
	stack := []rune{}
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}

	for _, char := range s {
		switch char {
		case '(', '[', '{':
			stack = append(stack, char)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[char] {
				return "NO"
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) == 0 {
		return "YES"
	}
	return "NO"
}

func main() {
	fmt.Println(isBalanced("{ [ ( ) ] }"))              // YES
	fmt.Println(isBalanced("{ [ ( ] ) }"))              // NO
	fmt.Println(isBalanced("{ ( ( [ ] ) [ ] ) [ ] }"))  // YES
}

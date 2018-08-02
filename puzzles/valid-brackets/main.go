package main

import "fmt"

func isValid(s string) bool {
	var stack []rune
	var popped rune

	for _, c := range s {
		if len(stack) == 0 {
			stack = append(stack, c)
		} else {
			if c == '{' || c == '[' || c == '(' {
				stack = append(stack, c)
			} else {
				popped = stack[len(stack)-1]
				if c == '}' && popped != '{' {
					return false
				} else if c == ']' && popped != '[' {
					return false
				} else if c == ')' && popped != '(' {
					return false
				} else {
					stack = stack[:len(stack)-1]
				}
			}
		}
	}

	if len(stack) > 0 {
		return false
	} else {
		return true
	}
}

func main() {
	var tcs = []string{"[]", "{}", "{{{()}}]", "[{}]()", "{[)", "{{", "(", "]", "()()"}
	for _, tc := range tcs {
		fmt.Println(tc)
		fmt.Println(isValid(tc))
	}
}

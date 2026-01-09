package main

import (
	"fmt"
	"strings"
)

func main() {
	var code string
	fmt.Scan(&code)

	stack := []string{}
	i := 0
	for i < len(code) {
		if strings.HasPrefix(code[i:], "bale") {
			stack = append(stack, "bale")
			i += 4
		} else if strings.HasPrefix(code[i:], "areh") {
			stack = append(stack, "areh")
			i += 4
		} else if strings.HasPrefix(code[i:], "kheir") {
			if len(stack) == 0 || stack[len(stack)-1] != "bale" {
				fmt.Println("NO")
				return
			}
			stack = stack[:len(stack)-1]
			i += 5
		} else if strings.HasPrefix(code[i:], "na") {
			if len(stack) == 0 || stack[len(stack)-1] != "areh" {
				fmt.Println("NO")
				return
			}
			stack = stack[:len(stack)-1]
			i += 2
		} else {
			i++
		}
	}

	if len(stack) == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

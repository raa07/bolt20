package main

import (
	"fmt"
)

func main() {
	var s string
	var t string

	s = "abc"
	t = "ahbgdc"

	fmt.Println(isSubsequence(s, t))
}

func isSubsequence(s string, t string) bool {
	//sArr := strings.Split(s, "")
	//tArr := strings.Split(t, "")

	return recursive(s, t)
}

func recursive(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	for i := 0; i < len(t); i++ {
		if s[0:1] == t[i:i+1] {
			return recursive(s[1:], t[i+1:])
		}
	}
	return false
}

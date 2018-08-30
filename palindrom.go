package main

import (
	"fmt"
	"os"
)

func isPalindrom(s string) bool {
	res := false
	if len(s) == 1 {
		res = true
	} else {
		if s[:1] == s[len(s)-1:len(s)] {
			if len(s)==2 {
				res = true
			} else {
				res = isPalindrom(s[1:len(s)-1])
			}
		}
	}
	return res
}

func main() {
	var palindrom string = os.Args[1]
	fmt.Println(palindrom + " is palindrom?")
	if isPalindrom(palindrom) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}


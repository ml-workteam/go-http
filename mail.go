package main

import (
	"fmt"
	"os"
)

func main(){
	/*
	 * usage: mail prefix str1 str2 str3
	 * no spaces?
	*/
	if len(os.Args) < 3 {
		os.Exit(1)
	} else {
		
		g := getThis(os.Args[2:], os.Args[1])
		for i,stri := range g {
			fmt.Printf("%d: %v\n", i, stri)
		}
	
	}
}

func getThis(s []string, prefix string) []string {
	var ret []string
	for _,stri := range s {
		if len(prefix) <= len(stri) && prefix == stri[:len(prefix)] {
			ret = append(ret, stri)
		}
	}
	return ret
}

package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(3)
	}
	res, err := http.Get(os.Args[1])
	if err != nil {
		os.Exit(1)
	}

	for mapi, hvar := range res.Header {
		fmt.Printf("%v: %v\n", mapi, hvar)
	}

}

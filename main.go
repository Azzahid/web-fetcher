package main

import (
	"fetcher/lib/fetcher"
	"fetcher/lib/localweb"
	"fmt"
	"os"
)

func main() {
	var args []string
	if len(os.Args) >= 2 {
		args = os.Args[1:]
	} else {
		panic("No variable received, return error")
	}

	if args[0] == "--metadata" {
		data := localweb.New(args[1])
		data.PrintMetadata()
	} else {
		fmt.Printf("Saving data for %v\n", args)
		fetcher.Fetch(args)
	}
}

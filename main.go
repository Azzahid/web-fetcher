package main

import (
	"fetcher/lib/web"
	"fmt"
	"os"
)

func main() {
	var args []string
	if len(os.Args) >= 2 {
		args = os.Args[1:]
	}

	if args[0] == "--metadata" {
		data := web.New(args[1])
		data.PrintMetadata()
	} else {
		fmt.Printf("Saving data for %v\n", args)
		web.Fetch(args)
	}
}

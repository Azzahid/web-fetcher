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
	// Init mounted directory
	mountedVol := os.Getenv("DOCKER_MOUNTED_VOL")
	if mountedVol != "" {
		err := os.MkdirAll(mountedVol, os.ModePerm)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
	}

	if args[0] == "--metadata" {
		data, err := localweb.Get(args[1])
		if err != nil {
			return
		}
		data.PrintMetadata()
	} else {
		fmt.Printf("Saving data for %v\n", args)
		fetcher.Fetch(args)
	}
}

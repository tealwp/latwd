package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	badUsageReturn = `
Usage: latwd <url> <max_depth> <max_breadth>

URL: [required] [string] - the base url to begin at
MAX_DEPTH: [not required] [integer] - the max depth to recursively call down to. Default is 5.
MAX_BREADTH: [not required] [integer] - the max breadth to recursively call from a single page. Default is 5.`
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("%s\n\n", badUsageReturn)
		os.Exit(0)
	}
	maxDepth := 5
	maxBreadth := 5
	if len(args) == 3 {
		var err error
		maxDepth, err = strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("maxDepth must be an integer")
			os.Exit(1)
		}
	}
	if len(args) == 4 {
		var err error
		maxBreadth, err = strconv.Atoi(args[3])
		if err != nil {
			fmt.Println("maxBreadth must be an integer")
			os.Exit(1)
		}
	}
	app := NewApp(args[1], maxDepth, maxBreadth)
	app.StartTrolling()
}

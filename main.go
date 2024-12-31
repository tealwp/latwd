package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	usageReturn = `
Usage: latwd <url> <max_depth> <max_breadth> <max_workers>

URL: [required] [string] - the base url to begin at
MAX_DEPTH: [not required] [integer] - the max depth to recursively call down to. Default is 5.
MAX_BREADTH: [not required] [integer] - the max breadth to recursively call from a single page. Default is 5.
MAX_WORKERS: [not required] [integer] - the max number of go routines used to run concurrently. Default is 5.`
)

// TODO: use proper flags, so users can conditionally use the defaults
func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("Incorrect number of arguments.\n %s\n\n", usageReturn)
		os.Exit(1)
	}
	if args[1] == "-h" || args[1] == "--help" {
		fmt.Printf("%s\n\n", usageReturn)
		os.Exit(0)
	}
	maxDepth := 5
	maxBreadth := 5
	maxWorkers := 5
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
	if len(args) == 5 {
		var err error
		maxWorkers, err = strconv.Atoi(args[4])
		if err != nil {
			fmt.Println("maxBreadth must be an integer")
			os.Exit(1)
		}
	}
	app := NewApp(args[1], maxDepth, maxBreadth, maxWorkers)
	defer app.Stop()

	app.StartTrolling()
	app.PrintTree()
}

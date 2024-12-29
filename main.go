package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: latwd <url!> <maxDepth?> <maxBreadth?>")
		os.Exit(1)
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

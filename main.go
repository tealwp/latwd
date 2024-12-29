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
	maxDepth := 10
	maxBreadth := 10
	if len(args) == 4 {
		var err error
		maxDepth, err = strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("maxDepth must be an integer")
			os.Exit(1)
		}
		maxBreadth, err = strconv.Atoi(args[3])
		if err != nil {
			fmt.Println("maxBreadth must be an integer")
			os.Exit(1)
		}
	}
	app := NewApp(maxDepth, maxBreadth)
	fmt.Println(app)
	app.StartTrolling(args[1])
	fmt.Printf("app links: %+v\n", app.links)
}

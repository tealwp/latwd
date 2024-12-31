package main

import (
	"fmt"
	"strings"
)

const (
	red        = "\033[31m"
	green      = "\033[32m"
	resetColor = "\033[0m"
)

type Link struct {
	URL      string
	Depth    int
	DeadLink DeadLink
	Children []*Link
	Parent   *Link
}

func (l *Link) printTree() {
	prefix := strings.Repeat("  ", l.Depth)
	var color string
	if l.DeadLink != nil {
		color = red
	} else {
		color = green
	}

	fmt.Printf("%s- %s%s%s\n", prefix, color, l.URL, resetColor)

	for _, child := range l.Children {
		child.printTree()
	}
}

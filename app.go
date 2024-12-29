package main

import (
	"errors"
	"fmt"
)

type Link struct {
	URL      string
	Depth    int
	DeadLink DeadLink
	Children []*Link
	Parent   *Link
	Client   *Client
}

type App struct {
	baseURL    string
	client     *Client
	maxDepth   int
	maxBreadth int
}

// input params should go in here and put onto the various fields
func NewApp(url string, maxDepth, maxBreadth int) *App {
	return &App{
		baseURL:    url,
		client:     NewClient(),
		maxDepth:   maxDepth,
		maxBreadth: maxBreadth,
	}
}

func (a *App) StartTrolling() {
	// recursively call Troll on the url, then its children, and so on and so forth
	link := Link{URL: a.baseURL, Depth: 1, Client: a.client}
	link.Troll(1, a.maxDepth, a.maxBreadth)
}

func (l *Link) Troll(currentDepth, maxDepth, maxBreadth int) {
	fmt.Printf("currentDepth: %d, url: %s\n", currentDepth, l.URL)
	if currentDepth >= maxDepth {
		return
	}
	// call link
	body, err := l.Client.Get(l.URL)
	// if deadlink, add to the link and return for next
	var deadLink DeadLink
	if err != nil && errors.As(err, &deadLink) {
		l.DeadLink = deadLink
		return
	}
	// else just print and return
	if err != nil {
		fmt.Printf("failed call for %s: %v\n", l.URL, err)
		return
	}
	childURLs, err := ParseHTML(body)
	if err != nil {
		fmt.Println(err)
		return
	}
	filteredLinks := FilterLinks(childURLs, l.URL)
	if len(filteredLinks) > maxBreadth {
		filteredLinks = filteredLinks[:maxBreadth]
	}
	for _, link := range filteredLinks {
		newL := &Link{URL: link, Depth: currentDepth + 1, Client: l.Client, Parent: l}
		l.Children = append(l.Children, newL)
		newL.Troll(currentDepth+1, maxDepth, maxBreadth)
	}
}

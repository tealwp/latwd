package main

import (
	"errors"
	"fmt"
	"sync"
)

// TODO: add workers/jobs to use concurrency for the http requests
type App struct {
	base       *Link
	client     *Client
	maxDepth   int
	maxBreadth int
	workers    []*Worker
	jobs       chan Job
	wg         sync.WaitGroup
}

// input params should go in here and put onto the various fields
func NewApp(url string, maxDepth, maxBreadth, numWorkers int) *App {
	app := &App{
		base:       &Link{URL: url, Depth: 1},
		client:     NewClient(),
		maxDepth:   maxDepth,
		maxBreadth: maxBreadth,
		workers:    make([]*Worker, numWorkers),
		jobs:       make(chan Job, 100), // Buffered channel for jobs
	}
	for i := 0; i < numWorkers; i++ {
		worker := NewWorker(i, app.jobs, &app.wg)
		app.workers[i] = worker
		app.wg.Add(1)
		go worker.Start()
	}
	return app
}

func (a *App) StartTrolling() {
	// recursively call Troll on the url, then its children, and so on and so forth
	a.troll(a.base, 1)
}

func (a *App) PrintTree() {
	a.base.printTree()
}

// AddJob adds a new job to the job queue.
func (app *App) AddJob(job Job) {
	app.jobs <- job
}

// Stop gracefully stops all workers.
func (app *App) Stop() {
	close(app.jobs) // Close the jobs channel to signal no more jobs
	app.wg.Wait()   // Wait for all workers to finish
}

func (a *App) troll(l *Link, currentDepth int) {
	// fmt.Printf("depth: %d, url: %s\n", currentDepth, l.URL)
	if currentDepth >= a.maxDepth {
		return
	}
	// call link
	body, err := a.client.Get(l.URL)
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
	childURLs, err := parseHTML(body)
	if err != nil {
		fmt.Println(err)
		return
	}
	filteredLinks := filterLinks(childURLs, l.URL)
	if len(filteredLinks) > a.maxBreadth {
		filteredLinks = filteredLinks[:a.maxBreadth]
	}
	for _, link := range filteredLinks {
		newL := &Link{URL: link, Depth: currentDepth + 1, Parent: l}
		l.Children = append(l.Children, newL)
		a.AddJob(
			Job{
				Function: func() {
					a.troll(newL, currentDepth+1)
				},
			},
		)
		a.troll(newL, currentDepth+1)
	}
}

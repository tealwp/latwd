package main

import "sync"

// Job represents a unit of work to be performed by a worker.
type Job struct {
	Function func()
}

// Worker represents a single worker that can execute jobs concurrently.
type Worker struct {
	ID       int
	JobQueue chan Job
	wg       *sync.WaitGroup
}

// NewWorker creates a new Worker instance.
func NewWorker(id int, jobQueue chan Job, wg *sync.WaitGroup) *Worker {
	return &Worker{
		ID:       id,
		JobQueue: jobQueue,
		wg:       wg,
	}
}

// Start starts the worker's execution loop.
func (w *Worker) Start() {
	defer w.wg.Done()
	for {
		select {
		case job := <-w.JobQueue:
			job.Function()
		default:
			return
		}
	}
}

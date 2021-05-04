package main

import "sync"

type Worker struct {
	stopped bool
	wg      *sync.WaitGroup
}

func (w *Worker) Stop() {
	w.stopped = true
	w.wg.Done()
}

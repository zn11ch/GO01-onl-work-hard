package main

import (
	"sync"
)

type CheckWorker struct {
	Worker
	words *Words
	quit  chan bool
	sm    *sync.Map
	once  *sync.Once
}

func NewCheckWorker(worker Worker, words *Words, quit chan bool, sm *sync.Map, once *sync.Once) *CheckWorker {
	return &CheckWorker{
		worker,
		words,
		quit,
		sm,
		once,
	}
}

func (c *CheckWorker) Work(words map[string]int) {

	for word, count := range words {
		value, ok := c.sm.LoadOrStore(word, count)
		if ok {
			c.sm.Store(word, value.(int)+count)

		}
		value1, _ := c.sm.Load(word)
		if value1.(int) >= 7 {
			c.once.Do(func() {
				close(c.quit)

			})
		}
	}
}

func (c *CheckWorker) Stop() {
	c.stopped = true

	c.wg.Done()
}

func (c *CheckWorker) eventLoop() {

	for {
		if c.stopped {
			return
		}
		select {
		case words, ok := <-c.words.C:

			if !ok {
				c.Stop()
				return
			} else {
				c.Work(words)
			}
		}
	}
}

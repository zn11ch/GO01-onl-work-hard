package workers

import (
	"fmt"
	"github.com/zn11ch/GO01-onl-work-hard/hw_lesson7/01/internal"
	"sync"
)

type CheckWorker struct {
	Worker
	words   *internal.Words
	quit    chan bool
	sm      *sync.Map
	once    *sync.Once
	checkWg *sync.WaitGroup
}

func NewCheckWorker(worker Worker, words *internal.Words, quit chan bool, sm *sync.Map, once *sync.Once, checkWg *sync.WaitGroup) *CheckWorker {
	return &CheckWorker{
		worker,
		words,
		quit,
		sm,
		once,
		checkWg,
	}
}

func (c *CheckWorker) Work(words map[string]int) {
	for word, count := range words {
		value, ok := c.sm.LoadOrStore(word, count)
		if ok {
			c.sm.Store(word, value.(int)+count)
		}
		checkValue, _ := c.sm.Load(word)
		if checkValue.(int) >= c.Config.CheckLimit {
			c.once.Do(func() {
				close(c.quit)
			})
			break
		}
	}
}

func (c *CheckWorker) Stop() {
	c.checkWg.Done()
	fmt.Println(c.checkWg)
	c.Stopped = true
	c.Wg.Done()
	c.checkWg.Wait()
}

func (c *CheckWorker) EventLoop() {

	for {
		if c.Stopped {
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

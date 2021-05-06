package workers

import (
	"github.com/zn11ch/GO01-onl-work-hard/hw_lesson7/01/internal"
	"strings"
	"sync"

	"unicode"
)

type CountWords struct {
	Worker
	line   *internal.Lines
	words  *internal.Words
	quit   chan bool
	sendWG *sync.WaitGroup
}

func NewCountWords(worker Worker, lines *internal.Lines, words *internal.Words, quit chan bool, sendWG *sync.WaitGroup) *CountWords {
	return &CountWords{
		Worker: worker,
		line:   lines,
		words:  words,
		quit:   quit,
		sendWG: sendWG,
	}
}

func (c *CountWords) Work(line string) {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}

	words := strings.FieldsFunc(line, f)
	m := make(map[string]int)
	for _, word := range words {
		if len(word) > 1 {
			m[word] += 1
		}
	}

	c.Write(m)

}

func (c *CountWords) Stop() {
	c.sendWG.Done()
	c.Stopped = true
	c.sendWG.Wait()
	c.words.SafeClose()
	c.Wg.Done()
}

func (c *CountWords) Write(data map[string]int) {
	c.words.C <- data
}

func (c *CountWords) EventLoop() {
	for {
		if c.Stopped {
			return
		}
		select {
		case line, ok := <-c.line.C:
			if ok {
				c.Work(line)
			} else {
				c.Stop()
			}
			//case <-c.quit:
			//	c.Stop()
			//	return
		}
	}
}

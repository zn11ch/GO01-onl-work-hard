package main

import (
	"fmt"
	"strings"
	"sync"

	"unicode"
)

type CountWords struct {
	Worker
	line   *Lines
	words  *Words
	quit   chan bool
	sendWG *sync.WaitGroup
}

func NewCountWords(worker Worker, lines *Lines, words *Words, quit chan bool, sendWG *sync.WaitGroup) *CountWords {
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

	c.sendWG.Add(1)
	defer c.sendWG.Done()
	c.Write(m)

}

func (c *CountWords) Stop() {

	c.stopped = true
	c.sendWG.Wait()
	c.words.SafeClose()
	c.wg.Done()
}

func (c *CountWords) Write(data map[string]int) {
	c.sendWG.Add(1)
	defer c.sendWG.Done()
	c.words.C <- data
}

func (c *CountWords) eventLoop() {

	for {
		if c.stopped {
			return
		}
		select {
		case line, ok := <-c.line.C:
			fmt.Println(line, ok)
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

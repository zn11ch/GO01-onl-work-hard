package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"sync"
)

type Lines struct {
	C    chan string
	once sync.Once
}

func NewLinesChannel() *Lines {
	return &Lines{C: make(chan string)}
}

func (mc *Lines) SafeClose() {
	mc.once.Do(func() {
		close(mc.C)
	})
}

type Words struct {
	C    chan map[string]int
	once sync.Once
}

func NewWordsChannel() *Words {
	return &Words{C: make(chan map[string]int)}
}

func (mc *Words) SafeClose() {
	mc.once.Do(func() {
		close(mc.C)
	})
}

func main() {
	absPath, _ := filepath.Abs("./hw_lesson7/01/test.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := NewLinesChannel()
	words := NewWordsChannel()
	quit := make(chan bool)
	sm := sync.Map{}
	once := sync.Once{}
	wg := sync.WaitGroup{}
	fileReaderWg := sync.WaitGroup{}
	countWordsWg := sync.WaitGroup{}
	n := 3

	for i := 0; i < n; i++ {
		wg.Add(3)
		w := Worker{false, &wg}
		r := NewFileReader(w, lines, quit, scanner, &fileReaderWg)
		c := NewCountWords(w, lines, words, quit, &countWordsWg)
		d := NewCheckWorker(w, words, quit, &sm, &once)
		go c.eventLoop()
		go r.eventLoop()
		go d.eventLoop()
	}

	wg.Wait()
}

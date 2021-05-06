package internal

import "sync"

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

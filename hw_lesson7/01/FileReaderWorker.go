package main

import (
	"bufio"
	"log"
	"sync"
)

type FileReader struct {
	Worker
	lines   *Lines
	quit    chan bool
	scanner *bufio.Scanner
	sendWG  *sync.WaitGroup
}

func NewFileReader(worker Worker, lines *Lines, quit chan bool, scanner *bufio.Scanner, sendWG *sync.WaitGroup) *FileReader {
	return &FileReader{
		Worker:  worker,
		lines:   lines,
		quit:    quit,
		scanner: scanner,
		sendWG:  sendWG,
	}
}

func (r *FileReader) Work() {

	for r.scanner.Scan() {
		r.Write(r.scanner.Text())
	}

	if err := r.scanner.Err(); err != nil {
		log.Fatal(err)
	}
	r.Stop()

}

func (r *FileReader) Write(data string) {
	r.sendWG.Add(1)
	defer r.sendWG.Done()
	r.lines.C <- data
}

func (r *FileReader) Stop() {
	r.stopped = true
	r.wg.Done()
	r.sendWG.Wait()
	r.lines.SafeClose()
}

func (r *FileReader) eventLoop() {
	for {
		if r.stopped {
			return
		}
		select {
		case <-r.quit:
			r.Stop()
			return
		default:
			r.Work()
		}
	}
}

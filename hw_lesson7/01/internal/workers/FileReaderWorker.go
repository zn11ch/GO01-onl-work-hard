package workers

import (
	"bufio"
	"github.com/zn11ch/GO01-onl-work-hard/hw_lesson7/01/internal"
	"log"
	"sync"
)

type FileReader struct {
	Worker
	lines   *internal.Lines
	quit    chan bool
	scanner *bufio.Scanner
	sendWG  *sync.WaitGroup
}

func NewFileReader(worker Worker, lines *internal.Lines, quit chan bool, scanner *bufio.Scanner, sendWG *sync.WaitGroup) *FileReader {
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
	r.lines.C <- data
}

func (r *FileReader) Stop() {
	r.Stopped = true
	r.Wg.Done()
	r.sendWG.Done()
	r.sendWG.Wait()
	r.lines.SafeClose()
}

func (r *FileReader) EventLoop() {

	for {
		if r.Stopped {
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

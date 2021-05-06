package workers

import (
	"github.com/zn11ch/GO01-onl-work-hard/hw_lesson7/01/internal"
	"sync"
)

type Worker struct {
	Stopped bool
	Wg      *sync.WaitGroup
	Config  *internal.Config
}

func (w *Worker) Stop() {
	w.Stopped = true
	w.Wg.Done()
}

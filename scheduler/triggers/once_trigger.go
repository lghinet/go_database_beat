package triggers

import (
	"time"
)

type OnceTrigger struct {
	done chan bool
}

func (trg *OnceTrigger) Tick(f func()) {
	trg.done = make(chan bool)
	go func(done chan bool) {
		f()
		done <- true
	}(trg.done)
}

func (trg *OnceTrigger) Cancel() {
	select {
	case <-trg.done:
		return
	case <-time.After(time.Minute):
		panic("timeout: cannot stop callback")
		return
	}
}

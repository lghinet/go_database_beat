package triggers

import (
	"time"
)

type TimeTrigger struct {
	Duration time.Duration
	ticker *time.Ticker
	canceled bool
}

func (trg *TimeTrigger) Tick(f func()) {
	trg.ticker = time.NewTicker(trg.Duration)
	go func() {
		for range trg.ticker.C {
			f()
		}
	}()
}

func (trg *TimeTrigger) Cancel() {
	trg.ticker.Stop()
	trg.canceled = true
}





func EveryMinute() *TimeTrigger {
	return &TimeTrigger{Duration:time.Minute}
}

func EveryFiveMinutes() *TimeTrigger{
	return &TimeTrigger{Duration:time.Minute*5}
}

func EveryHour() *TimeTrigger {
	return &TimeTrigger{Duration:time.Hour}
}

func Every30Seconds() *TimeTrigger {
	return &TimeTrigger{Duration:time.Second*30}
}

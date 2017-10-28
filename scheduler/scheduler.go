package scheduler

import "log"

type Scheduler struct {
	Jobs []Job
}

func (scheduler *Scheduler) Start() {
	for _, job := range scheduler.Jobs {
		func(_job Job) {
			_job.GetTrigger().Tick(func() {
				_job.Run()
				log.Println("tick job ", _job.GetName())
			})
		}(job)
	}
}

func (scheduler *Scheduler) Stop() {
	for _, job := range scheduler.Jobs {
		job.GetTrigger().Cancel()
	}
}

func (scheduler *Scheduler) RegisterJob(job Job) {
	scheduler.Jobs = append(scheduler.Jobs, job)
}

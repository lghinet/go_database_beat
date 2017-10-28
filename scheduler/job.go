package scheduler

type Job interface {
	Run()
	GetTrigger() Trigger
	GetName() string
}

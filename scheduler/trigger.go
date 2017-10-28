package scheduler

type Trigger interface{
	Tick(callback func())
	Cancel()
}

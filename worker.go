package routinepool

type worker struct {
	taskChan chan func()

	lastSheduleTime int64
}

func NewWorker() *worker {
	return &worker{
		taskChan: make(chan func()),
	}
}

func (w worker) Start() {
	for f := range w.taskChan {
		if f == nil {
			return
		}

		f()
	}
}

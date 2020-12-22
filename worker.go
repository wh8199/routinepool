package routinepool

type worker struct {
	taskChan chan func()

	updateTime int64
}

func (w worker) Start() {
	for f := range w.taskChan {
		if f == nil {
			return
		}

		f()
	}
}

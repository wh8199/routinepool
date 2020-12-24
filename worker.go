package routinepool

type worker struct {
	taskChan chan func()

	lastSheduleTime int64

	pool *RoutinePool
}

func NewWorker(pool *RoutinePool) *worker {
	return &worker{
		taskChan: make(chan func()),
		pool:     pool,
	}
}

func (w *worker) Start() {
	for f := range w.taskChan {
		if f == nil {
			return
		}

		f()

		w.pool.Recycle(w)
	}
}

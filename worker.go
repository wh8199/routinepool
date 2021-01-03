package routinepool

import (
	"time"
)

type worker struct {
	taskChan        chan func()
	lastSheduleTime time.Time
	//pool            *RoutinePool
}

func NewWorker(pool *RoutinePool) *worker {
	return &worker{
		taskChan: make(chan func()),
		//pool:            pool,
		lastSheduleTime: time.Now(),
	}
}

func (w *worker) Stop() {
	close(w.taskChan)
}

func (w *worker) Start() {
	for {
		select {
		case f := <-w.taskChan:
			f()
			//w.pool.Recycle(w)
		}
	}
}

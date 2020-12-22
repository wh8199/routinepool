package routinepool

import "sync"

type RoutinePool struct {
	Config *RoutinePoolConfig

	WorkerNumber int64 `json:"workerNumber"`
	Lock         sync.Mutex

	ReadyWorkers []*worker
}

func NewRoutinePool(config *RoutinePoolConfig) *RoutinePool {
	return &RoutinePool{
		Config:       config,
		WorkerNumber: 0,
		Lock:         sync.Mutex{},
	}
}

func (r *RoutinePool) getReadyWorker() *worker {
	r.Lock.Lock()
	defer r.Lock.Unlock()

	if len(r.ReadyWorkers) > 0 {
		w := r.ReadyWorkers[len(r.ReadyWorkers)-1]
		r.ReadyWorkers = r.ReadyWorkers[:len(r.ReadyWorkers)-1]
		return w
	}

	r.WorkerNumber++
	w := NewWorker()
	go w.Start()

	return w
}

func (r RoutinePool) SubmitWorker(f func()) {
	w := r.getReadyWorker()

	w.taskChan <- f
}

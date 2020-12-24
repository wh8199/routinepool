package routinepool

import (
	"fmt"
	"sync"
	"time"
)

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

func (r *RoutinePool) Start() {
	for {
		fmt.Println(r.WorkerNumber)

		time.Sleep(5 * time.Second)
	}
}

func (r *RoutinePool) getReadyWorker() *worker {
	r.Lock.Lock()
	defer r.Lock.Unlock()

	r.WorkerNumber = r.WorkerNumber + 1

	if len(r.ReadyWorkers) > 0 {
		w := r.ReadyWorkers[len(r.ReadyWorkers)-1]
		r.ReadyWorkers = r.ReadyWorkers[:len(r.ReadyWorkers)-1]
		return w
	}

	w := NewWorker(r)
	go w.Start()

	return w
}

func (r *RoutinePool) SubmitWorker(f func()) {
	w := r.getReadyWorker()

	w.taskChan <- f
}

func (r *RoutinePool) Recycle(w *worker) {
	w.lastSheduleTime = time.Now().Unix()

	r.Lock.Lock()
	defer r.Lock.Unlock()

	r.ReadyWorkers = append(r.ReadyWorkers, w)
	r.WorkerNumber = r.WorkerNumber - 1
}

package routinepool

import (
	"sync"
	"time"

	"github.com/wh8199/log"
)

type RoutinePool struct {
	Config *RoutinePoolConfig

	WorkerNumber int64 `json:"workerNumber"`
	Lock         sync.Mutex

	Log          log.LoggingInterface
	ReadyWorkers []*worker
}

func NewRoutinePool(config *RoutinePoolConfig) *RoutinePool {
	return &RoutinePool{
		Config:       config,
		WorkerNumber: 0,
		Lock:         sync.Mutex{},
		Log:          log.NewLogging("routinepool", config.LogLevel, 2),
	}
}

func (r *RoutinePool) Start() {
	r.StartCleanWorkers()
}

func (r *RoutinePool) cleanWorkerOnce() {
	r.Lock.Lock()
	defer r.Lock.Unlock()

	if len(r.ReadyWorkers) == 0 {
		return
	}

	cutIndex := -1
	now := time.Now()
	for index, worker := range r.ReadyWorkers {
		if now.After(worker.lastSheduleTime.Add(r.Config.MaxIdleTime)) {
			cutIndex = index
			break
		}
	}

	if cutIndex == -1 {
		return
	}

	if len(r.ReadyWorkers) == 1 {
		r.ReadyWorkers = []*worker{}
	} else {
		r.ReadyWorkers = r.ReadyWorkers[cutIndex+1:]
	}
}

func (r *RoutinePool) StartCleanWorkers() {
	ticker := time.NewTicker(r.Config.CleanInterval)

	for {
		select {
		case <-ticker.C:
			r.Log.Debug("Start clean workers")
			r.cleanWorkerOnce()
			r.Log.Debug("Clean worker done")
		}
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
	w.lastSheduleTime = time.Now()

	r.Lock.Lock()
	defer r.Lock.Unlock()

	r.ReadyWorkers = append(r.ReadyWorkers, w)
	r.WorkerNumber = r.WorkerNumber - 1
}

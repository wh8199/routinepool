package routinepool

import (
	"sync"
	"time"

	"github.com/wh8199/log"
)

type RoutinePool struct {
	config       *RoutinePoolConfig
	workerNumber int64
	lock         sync.Mutex
	log          log.LoggingInterface
	readyWorkers []*worker
}

func NewRoutinePool(config *RoutinePoolConfig) *RoutinePool {
	return &RoutinePool{
		config:       config,
		workerNumber: 0,
		lock:         sync.Mutex{},
		log:          log.NewLogging("routinepool", config.LogLevel, 2),
	}
}

func (r *RoutinePool) Start() {
	go r.StartCleanWorkers()

	for {
		r.log.Debugf("Current worker number: %d", r.workerNumber)
		r.log.Debugf("Current number of ready worker: %d", len(r.readyWorkers))
		time.Sleep(10 * time.Second)
	}
}

func (r *RoutinePool) cleanWorkerOnce() {
	r.lock.Lock()
	defer r.lock.Unlock()

	if len(r.readyWorkers) == 0 {
		return
	}

	cutIndex := -1
	now := time.Now()
	for index, worker := range r.readyWorkers {
		if now.After(worker.lastSheduleTime.Add(r.config.MaxIdleTime)) {
			cutIndex = index
			break
		}
	}

	if cutIndex == -1 {
		return
	}

	if len(r.readyWorkers) == 1 {
		r.readyWorkers = []*worker{}
	} else {
		r.readyWorkers = r.readyWorkers[cutIndex+1:]
	}
}

func (r *RoutinePool) StartCleanWorkers() {
	ticker := time.NewTicker(r.config.CleanInterval)

	for {
		select {
		case <-ticker.C:
			r.log.Debug("Start clean workers")
			r.cleanWorkerOnce()
			r.log.Debug("Clean worker done")
		}
	}
}

func (r *RoutinePool) getReadyWorker() *worker {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.workerNumber = r.workerNumber + 1

	if len(r.readyWorkers) > 0 {
		w := r.readyWorkers[len(r.readyWorkers)-1]
		r.readyWorkers = r.readyWorkers[:len(r.readyWorkers)-1]
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

	r.lock.Lock()
	defer r.lock.Unlock()

	r.readyWorkers = append(r.readyWorkers, w)
	r.workerNumber = r.workerNumber - 1
}

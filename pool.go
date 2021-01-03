package routinepool

import (
	"sync"

	"github.com/wh8199/log"
)

type RoutinePool struct {
	config       *RoutinePoolConfig
	workerNumber uint64
	lock         sync.RWMutex
	log          log.LoggingInterface
	readyWorker  uint64
	workerChan   chan func()
}

func NewRoutinePool(config *RoutinePoolConfig) *RoutinePool {
	routinePool := &RoutinePool{
		config:       config,
		workerNumber: 0,
		lock:         sync.RWMutex{},
		log:          log.NewLogging("routinepool", config.LogLevel, 2),
		readyWorker:  0,
		workerChan:   make(chan func(), 128),
	}

	return routinePool
}

func (r *RoutinePool) SubmitWorker(f func()) {
	createWorker := false
	r.lock.Lock()
	if r.readyWorker <= 0 {
		createWorker = true
	}
	r.lock.Unlock()

	if createWorker {
		go r.startWorker()
		r.workerChan <- f
		return
	}

	r.lock.Lock()
	r.workerChan <- f
	r.readyWorker--
	r.lock.Unlock()
}

func (r *RoutinePool) donewWorker() {
	r.lock.Lock()
	r.readyWorker++
	r.lock.Unlock()
}

func (r *RoutinePool) startWorker() {
	for {
		select {
		case f := <-r.workerChan:
			f()
			r.donewWorker()
		}
	}
}

package routinepool

import (
	"fmt"
	"log"
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
	go r.CleanWorker()

	for {
		fmt.Println(r.WorkerNumber)
		fmt.Println(len(r.ReadyWorkers))
		time.Sleep(5 * time.Second)
	}
}

func (r *RoutinePool) CleanWorker() {
	ticker := time.NewTicker(r.Config.CleanInterval)

	for {
		select {
		case <-ticker.C:
			log.Println("Clean workers")
			r.Lock.Lock()

			if len(r.ReadyWorkers) == 0 {
				r.Lock.Unlock()
				continue
			}

			cutIndex := -1
			now := time.Now()
			for index, worker := range r.ReadyWorkers {
				if now.After(worker.lastSheduleTime.Add(r.Config.MaxIdleTime)) {
					cutIndex = index
					break
				}
			}

			log.Println(cutIndex)

			if cutIndex == -1 {
				r.Lock.Unlock()
				continue
			}

			if len(r.ReadyWorkers) == 1 {
				r.ReadyWorkers = []*worker{}
			} else {
				r.ReadyWorkers = r.ReadyWorkers[cutIndex+1:]
			}

			r.Lock.Unlock()
			log.Println("Clean workers done")
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

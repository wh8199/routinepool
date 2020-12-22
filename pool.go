package routinepool

import "sync"

type RoutinePool struct {
	Config *RoutinePoolConfig

	WorkerNumber int64 `json:"workerNumber"`
	Lock         sync.Mutex

	Workers []*worker
}

func NewRoutinePool(config *RoutinePoolConfig) *RoutinePool {
	return &RoutinePool{
		Config:       config,
		WorkerNumber: 0,
		Lock:         sync.Mutex{},
	}
}

func (r RoutinePool) getReadyWorker() *worker {
	return nil
}

func (r RoutinePool) SubmitWorker(f func()) {
	//go f()
}

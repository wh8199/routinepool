package routinepool

import (
	"testing"
)

func BenchmarkCreatingWorker(b *testing.B) {
	config := DefaultRouterPoolConfig()
	pool := NewRoutinePool(config)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		pool.getReadyWorkerWithoutPool()
	}
}

func BenchmarkCreatingWorkerWithPool(b *testing.B) {
	config := DefaultRouterPoolConfig()
	pool := NewRoutinePool(config)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		w := pool.getReadyWorker()
		pool.workerPool.Put(w)
	}
}

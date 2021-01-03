package routinepool

/*
func BenchmarkCreatingWorker(b *testing.B) {
	config := DefaultRouterPoolConfig()
	pool := NewRoutinePool(config)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		pool.getReadyWorkerWithoutPool()
	}

	b.StopTimer()
}

func BenchmarkCreatingWorkerWithPool(b *testing.B) {
	config := DefaultRouterPoolConfig()
	pool := NewRoutinePool(config)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		w := pool.getReadyWorker()
		pool.workerPool.Put(w)
	}

	b.StopTimer()
}

func BenchmarkNoPool(b *testing.B) {
	var wg sync.WaitGroup
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			demoFunc()
			wg.Done()
		}()
	}

	wg.Wait()

	b.StopTimer()
}

func BenchmarkPool(b *testing.B) {
	config := DefaultRouterPoolConfig()
	config.WithCleanWorkerInterval("60s")
	config.WithMaxIdleTime("120s")
	pool := NewRoutinePool(config)
	b.ResetTimer()
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)

		pool.SubmitWorker(func() {
			demoFunc()
			wg.Done()
		})
	}

	wg.Wait()
	b.StopTimer()
}
*/

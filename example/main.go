package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/wh8199/log"
	"github.com/wh8199/routinepool"

	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		http.ListenAndServe("0.0.0.0:6061", nil)
	}()

	config := routinepool.DefaultRouterPoolConfig()
	config.WithCleanWorkerInterval("10s")
	config.WithMaxIdleTime("20s")
	config.WithLogLevel(log.DEBUG_LEVEL)
	pool := routinepool.NewRoutinePool(config)
	go pool.StartCleanWorkers()

	count := 100000000
	actualCount := 0
	actualCountLock := sync.Mutex{}

	f := func() {
		time.Sleep(time.Duration(10) * time.Millisecond)
		actualCountLock.Lock()
		actualCount++
		actualCountLock.Unlock()
	}

	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		wg.Add(1)
		pool.SubmitWorker(func() {
			f()
			wg.Done()
		})
	}

	wg.Wait()

	fmt.Printf("Actual count is %d\n", actualCount)
}

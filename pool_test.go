package routinepool

import (
	"runtime"
	"sync"
	"testing"
	"time"

	"github.com/wh8199/log"
)

var curMem uint64

const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
)

const (
	n = 1000000
)

func demoFunc() {
	time.Sleep(time.Duration(10) * time.Microsecond)
}

func TestNoPool(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			demoFunc()
			wg.Done()
		}()
	}

	wg.Wait()
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	curMem = mem.TotalAlloc/MiB - curMem
	t.Logf("memory usage:%d MB", curMem)
}

func TestPool(t *testing.T) {
	config := DefaultRouterPoolConfig()
	config.WithCleanWorkerInterval("60s")
	config.WithMaxIdleTime("120s")
	config.WithLogLevel(log.DEBUG_LEVEL)
	pool := NewRoutinePool(config)

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		pool.SubmitWorker(func() {
			demoFunc()
			wg.Done()
		})
	}

	wg.Wait()
	t.Log(pool.readyWorker)

	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	curMem = mem.TotalAlloc/MiB - curMem
	t.Logf("memory usage:%d MB", curMem)
}

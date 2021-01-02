package main

import (
	"sync"
	"time"

	"github.com/wh8199/log"
	"github.com/wh8199/routinepool"

	"net/http"
	_ "net/http/pprof"
)

const (
	RunTimes   = 100000000
	BenchParam = 10
)

var curMem uint64

const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	// GiB // 1073741824
	// TiB // 1099511627776             (超过了int32的范围)
	// PiB // 1125899906842624
	// EiB // 1152921504606846976
	// ZiB // 1180591620717411303424    (超过了int64的范围)
	// YiB // 1208925819614629174706176
)

const (
	Param    = 100
	AntsSize = 1000
	TestSize = 10000
	n        = 100000000
)

func demoFunc() {
	time.Sleep(time.Duration(BenchParam) * time.Millisecond)
}

func main() {
	go func() {
		http.ListenAndServe("0.0.0.0:6061", nil)
	}()

	config := routinepool.DefaultRouterPoolConfig()
	config.WithCleanWorkerInterval("10s")
	config.WithMaxIdleTime("20s")
	config.WithLogLevel(log.DEBUG_LEVEL)
	pool := routinepool.NewRoutinePool(config)

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		pool.SubmitWorker(func() {
			demoFunc()
			wg.Done()
		})
	}

	wg.Wait()
}

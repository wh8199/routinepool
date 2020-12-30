package main

import (
	"fmt"
	"time"

	"github.com/wh8199/log"
	"github.com/wh8199/routinepool"
)

func main() {
	config := routinepool.DefaultRouterPoolConfig()
	config.WithCleanWorkerInterval("10s")
	config.WithMaxIdleTime("20s")
	config.WithLogLevel(log.DEBUG_LEVEL)
	pool := routinepool.NewRoutinePool(config)

	pool.SubmitWorker(func() {
		time.Sleep(1 * time.Second)
		fmt.Println("worker1")
	})

	pool.SubmitWorker(func() {
		time.Sleep(2 * time.Second)
		fmt.Println("worker2")
	})

	pool.Start()
}

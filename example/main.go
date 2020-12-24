package main

import (
	"fmt"
	"time"

	"github.com/wh8199/routinepool"
)

func main() {
	config := routinepool.DefaultRouterPoolConfig()
	pool := routinepool.NewRoutinePool(config)

	pool.SubmitWorker(func() {
		time.Sleep(1 * time.Second)
		fmt.Println("sss")
	})

	pool.Start()
}

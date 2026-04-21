package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"

	"go-throughput-lab/internal/baseline"
)

func run(name string, fn func(int, int) baseline.Result, workers, events int) {
	start := time.Now()

	result := fn(workers, events)

	elapsed := time.Since(start)
	throughput := float64(result.Processed) / elapsed.Seconds()

	fmt.Println("===", name, "===")
	fmt.Printf("Workers: %d\n", workers)
	fmt.Printf("Total Events: %d\n", events)
	fmt.Printf("Processed: %d\n", result.Processed)
	fmt.Printf("Elapsed: %s\n", elapsed)
	fmt.Printf("Throughput: %.2f events/sec\n\n", throughput)
}

func main() {
	runtime.SetBlockProfileRate(1)

	go func() {
		fmt.Println("pprof running on :6060")
		http.ListenAndServe(":6060", nil)
	}()

	numWorkers := runtime.NumCPU() * 2
	totalEvents := 20_000_000

	run("Single Channel", baseline.RunSingleChannel, numWorkers, totalEvents)

	run("Sharded Channels", baseline.RunShardedChannels, numWorkers, totalEvents)

	run(
		"Sharded + Multi Producer",
		func(w, e int) baseline.Result {
			return baseline.RunShardedChannelsMultiProducer(w, e, runtime.NumCPU())
		},
		numWorkers,
		totalEvents,
	)

	run(
		"Sharded + Multi Producer + Batch",
		func(w, e int) baseline.Result {
			return baseline.RunShardedChannelsWithBatch(
				w,
				e,
				runtime.NumCPU(),
				64,
			)
		},
		numWorkers,
		totalEvents,
	)

	// manter vivo sem deadlock
	time.Sleep(30 * time.Second)
}
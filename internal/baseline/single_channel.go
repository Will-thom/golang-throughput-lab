package baseline

import (
	"sync"
	"sync/atomic"
)

type Event struct {
	ID int64
}

type Result struct {
	Processed int64
}

func RunSingleChannel(numWorkers int, totalEvents int) Result {
	var processed int64

	ch := make(chan Event, 1024) // buffer pequeno → força contenção

	var wg sync.WaitGroup

	// Workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range ch {
				// simulação mínima de trabalho (CPU-bound leve)
				atomic.AddInt64(&processed, 1)
			}
		}()
	}

	// Producer
	for i := 0; i < totalEvents; i++ {
		ch <- Event{ID: int64(i)}
	}

	close(ch)
	wg.Wait()

	return Result{
		Processed: processed,
	}
}
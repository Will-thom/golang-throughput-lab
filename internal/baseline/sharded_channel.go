package baseline

import (
	"sync"
	"sync/atomic"
)

func RunShardedChannels(numWorkers int, totalEvents int) Result {
	var processed int64

	channels := make([]chan Event, numWorkers)
	var wg sync.WaitGroup

	// Criar um channel por worker
	for i := 0; i < numWorkers; i++ {
		channels[i] = make(chan Event, 1024)

		wg.Add(1)
		go func(ch chan Event) {
			defer wg.Done()
			for range ch {
				atomic.AddInt64(&processed, 1)
			}
		}(channels[i])
	}

	// Distribuição simples: round-robin
	for i := 0; i < totalEvents; i++ {
		idx := i % numWorkers
		channels[idx] <- Event{ID: int64(i)}
	}

	// Fechar todos os channels
	for _, ch := range channels {
		close(ch)
	}

	wg.Wait()

	return Result{
		Processed: processed,
	}
}
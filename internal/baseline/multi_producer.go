package baseline

import (
	"sync"
	"sync/atomic"
)

func RunShardedChannelsMultiProducer(numWorkers int, totalEvents int, numProducers int) Result {
	var processed int64

	channels := make([]chan Event, numWorkers)
	var wg sync.WaitGroup

	// Workers
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

	// Producers
	var producerWg sync.WaitGroup
	eventsPerProducer := totalEvents / numProducers

	for p := 0; p < numProducers; p++ {
		start := p * eventsPerProducer
		end := start + eventsPerProducer

		producerWg.Add(1)
		go func(start, end int) {
			defer producerWg.Done()

			for i := start; i < end; i++ {
				idx := i % numWorkers
				channels[idx] <- Event{ID: int64(i)}
			}
		}(start, end)
	}

	producerWg.Wait()

	// Fechar channels
	for _, ch := range channels {
		close(ch)
	}

	wg.Wait()

	return Result{
		Processed: processed,
	}
}
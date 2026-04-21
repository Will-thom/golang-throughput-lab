package baseline

import (
	"sync"
	"sync/atomic"
)

func RunShardedChannelsWithBatch(
	numWorkers int,
	totalEvents int,
	numProducers int,
	batchSize int,
) Result {

	var processed int64

	channels := make([]chan []Event, numWorkers)
	var wg sync.WaitGroup

	// Workers
	for i := 0; i < numWorkers; i++ {
		channels[i] = make(chan []Event, 1024)

		wg.Add(1)
		go func(ch chan []Event) {
			defer wg.Done()
			for batch := range ch {
				for range batch {
					atomic.AddInt64(&processed, 1)
				}
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

			batch := make([]Event, 0, batchSize)

			for i := start; i < end; i++ {
				batch = append(batch, Event{ID: int64(i)})

				if len(batch) == batchSize {
					idx := i % numWorkers
					channels[idx] <- batch

					// reutiliza slice (evita alocação)
					batch = batch[:0]
				}
			}

			// flush final
			if len(batch) > 0 {
				idx := end % numWorkers
				channels[idx] <- batch
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
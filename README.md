# Go Throughput Lab

## Problem

High-throughput event processing systems often suffer from contention, excessive allocations, and poor scalability when using naive concurrency models (e.g. single shared channels or unbounded goroutines).

This project explores different concurrency strategies in Go to maximize throughput while keeping resource usage predictable.

---

## Why this exists

In real-world systems (stream processing, ingestion pipelines, async workloads), it's common to hit bottlenecks caused by:

- shared channel contention
- scheduler pressure from too many goroutines
- memory allocation overhead in hot paths

This repository is a hands-on lab to experiment, measure, and understand these trade-offs.

---

## Approaches explored

### 1. Single Channel Worker Pool
- Simple design
- Centralized queue
- Easy to reason about

**Limitations:**
- High contention under load
- Throughput degrades with more workers

---

### 2. Sharded Channels (Work Distribution)
- Events distributed across multiple channels
- Reduces contention
- Improves parallelism

**Trade-offs:**
- Increased complexity
- Potential imbalance between shards

---

### 3. Allocation Reduction Strategies
- Reuse of objects
- Avoid unnecessary memory allocations in hot paths

**Impact:**
- Lower GC pressure
- More stable latency

---

## Benchmark (example)

| Strategy            | Events     | Time       | Throughput        |
|--------------------|-----------|------------|-------------------|
| Single Channel     | 20,000,000| ~4.05s     | ~4.9M events/sec  |
| Sharded Channels   | 20,000,000| ~4.44s     | ~4.5M events/sec  |

> Note: Results may vary depending on CPU and runtime conditions.

---

## Key Insights

- Simpler architectures are not always slower — contention patterns matter more than structure
- Increasing workers does not guarantee better throughput
- Memory allocation patterns significantly impact performance in high-throughput systems

---

## Trade-offs & Design Decisions

- Prioritized throughput over simplicity in some experiments
- Avoided external dependencies to keep focus on Go runtime behavior
- Explicitly chose not to use advanced frameworks to isolate core concurrency patterns

---

## What I would improve next

- Introduce dynamic worker scaling
- Add observability (metrics, tracing)
- Simulate real-world workloads (I/O, backpressure)
- Explore lock-free data structures

---

## How to run

```bash
go run ./cmd/app

## Final note

This is not a production-ready system.

It is a focused exploration of concurrency trade-offs in Go for high-throughput workloads.

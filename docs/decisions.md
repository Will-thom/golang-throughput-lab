# Design Decisions

## Why not a single channel?

While a single channel simplifies coordination, it becomes a bottleneck under high contention scenarios.

## Why sharding?

Sharding reduces contention by distributing load, at the cost of increased complexity and potential imbalance.

## Why not use existing frameworks?

The goal is to understand runtime behavior, not abstract it away.

## Trade-off philosophy

This project prioritizes:
- throughput
- predictability

Over:
- simplicity
- abstraction

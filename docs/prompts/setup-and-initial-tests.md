---

# 🚀 PASSO 1 — Bootstrap + Baseline (single channel → alto contention)

## 🎯 Objetivo deste passo

Construir a versão mais simples possível:

* 1 channel global
* múltiplos workers consumindo dele
* gerar carga sintética
* medir throughput

👉 Isso vai criar **contenção real** → exatamente o que queremos observar depois.

---

## 📁 Estrutura inicial

```
go-throughput-lab/
 ├── cmd/
 │    └── app/
 │         └── main.go
 ├── internal/
 │    └── baseline/
 │         └── single_channel.go
 ├── go.mod
```

---

## 📦 Arquivo 1

### 📍 `go.mod`

```go
module go-throughput-lab

go 1.22
```

---

## 📦 Arquivo 2 (core do baseline)

### 📍 `internal/baseline/single_channel.go`

```go
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
```

---

## 📦 Arquivo 3 (runner)

### 📍 `cmd/app/main.go`

```go
package main

import (
	"fmt"
	"runtime"
	"time"

	"go-throughput-lab/internal/baseline"
)

func main() {
	numWorkers := runtime.NumCPU() * 2
	totalEvents := 5_000_000

	start := time.Now()

	result := baseline.RunSingleChannel(numWorkers, totalEvents)

	elapsed := time.Since(start)

	throughput := float64(result.Processed) / elapsed.Seconds()

	fmt.Println("=== Single Channel Baseline ===")
	fmt.Printf("Workers: %d\n", numWorkers)
	fmt.Printf("Total Events: %d\n", totalEvents)
	fmt.Printf("Processed: %d\n", result.Processed)
	fmt.Printf("Elapsed: %s\n", elapsed)
	fmt.Printf("Throughput: %.2f events/sec\n", throughput)
}
```

---

## ▶️ Como rodar

Na raiz do projeto:

```bash
go mod tidy
go run ./cmd/app
```

---

## ✅ O que você deve observar

Não precisa otimizar nada agora. Só observe:

* throughput final
* uso de CPU (se quiser rodar com `top` ou similar)
* tempo total

👉 Esse resultado é sua **linha de base**

---

## 🧠 O que está acontecendo (importante)

Esse design tem:

* **1 único channel compartilhado**
* múltiplos workers competindo por ele
* sincronização interna do channel

➡️ Isso gera:

* contenção
* possível blocking
* limite de escala

👉 Exatamente o problema que vamos resolver depois.

---

## ✅ Critério de conclusão deste passo

Antes de continuar, você precisa:

* rodar o programa
* ver o throughput
* confirmar que tudo funciona


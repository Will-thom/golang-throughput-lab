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
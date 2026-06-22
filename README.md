# 🏢 MaxInTec Digital Twin API

A high-performance, resilient Digital Twin backend engine built in **Go (Golang)** using the **Gin Gonic** framework. This platform acts as a centralized data orchestration hub, integrating corporate database landscapes with real-time telemetric streams.

---

## 🏗️ Architectural Overview (Package by Feature)

The project is structured around **Domain Capabilities (Features)**. Each package inside the core directory is self-contained, encapsulating its own routing handlers, business logic, entities, and data access adapters to prevent cyclic dependencies.

* **`cmd/main.go`**: Application entry point and structural dependency injection.
* **`internal/config/`**: Global environment variables and system setup parsing.
* **`internal/ordemservico/`**: Service Orders domain handling HTTP requests, business logic, and dedicated MSSQL queries.
* **`internal/pessoa/`**: Client, technician, and user records management containing specific database integration.
* **`internal/platform/database/`**: Centralized Microsoft SQL Server connection pooling used across packages.
* **`internal/rastreador/`**: Core telemetry stream engine grouping the worker loop, API client, and live RAM cache.

---

## ⚡ Concurrency & Performance Engine (`internal/rastreador`)

### 🛰️ Background Worker (`worker.go`)
Spawns as an isolated **Goroutine** upon application startup. It continuously handles incoming telemetric streams from tracking vendor APIs completely detached from the main Gin HTTP thread.

### 🧠 In-Memory Cache (`cache.go`)
To avoid choking disk-based production databases with heavy telemetry writes, the background worker stores real-time location metrics directly in a thread-safe RAM cache layer.

### 📈 Operational Impact
API request handlers bypass relational storage entirely for live location tracking. This design achieves sub-millisecond delivery to monitoring clients with **0% write-overhead** on the legacy production infrastructure.

---

## 🛠️ Technology Stack

* **Language Engine:** Go (Golang)
* **HTTP Router & REST API:** Gin Gonic Framework
* **Production Database:** Microsoft SQL Server (MSSQL) via database platform pooling
* **Concurrency Architecture:** Go Background Workers coupled with an In-Memory Thread-Safe Cache Engine

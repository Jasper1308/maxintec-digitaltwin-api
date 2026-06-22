🏢 MaxInTec Digital Twin API
A high-performance, resilient Digital Twin backend engine built in Go (Golang) using the Gin Gonic framework. This platform acts as a centralized data orchestration hub, integrating corporate database landscapes with real-time telemetric streams.

🏗️ Architectural Overview (Package by Feature)
The project is structured around Domain Capabilities (Features). Each package inside internal/ is self-contained, encapsulating its own routing handlers, business logic, entities, and data access adapters to prevent cyclic dependencies.

cmd/main.go: Application entry point & dependency injection

internal/config/: Environment variables & system setup parsing

internal/ordemservico/: Service Orders domain (handlers, business logic & mssql)

internal/pessoa/: Users, technicians or clients domain (handlers, business logic & mssql)

internal/platform/database/: Centralized MSSQL connection pooling

internal/rastreador/: Telemetry stream intake engine (worker, client & RAM cache)

⚡ Concurrency & Performance Engine (internal/rastreador)
Background Worker (worker.go): Spawns as an isolated Goroutine upon application startup. It continuously handles incoming telemetric streams from tracking vendor APIs completely detached from the main Gin HTTP thread.

In-Memory Cache (cache.go): To avoid choking disk-based production databases with telemetry writes, the worker stores real-time location metrics directly in a thread-safe RAM cache layer.

Impact: API request handlers bypass relational storage for live location tracking, achieving sub-millisecond delivery with 0% overhead on the legacy infrastructure.

🛠️ Technology Stack
Language: Go (Golang)

HTTP Router / REST API: Gin Gonic

Database Platform: Microsoft SQL Server (MSSQL) via platform/database

Concurrency Core: Go Background Workers & In-Memory Thread-Safe Cache

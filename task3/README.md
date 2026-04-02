# Task 3: Skeleton A of a future full-fledged microservice application

A distributed system for fetching and processing information about GitHub repositories. This project demonstrates a complete microservices workflow, evolving from a simple CLI tool into a robust architecture with an API Gateway, Processor, and Collector services.

## Architecture

The system is built with **Clean Architecture** principles in mind, ensuring high maintainability and testability. All internal communication between services is strictly performed via gRPC.

The project consists of 4 specialized microservices:

### 1. API Gateway Service

- **Role**: Entry point for external clients.
- **Transports**: REST (external) and gRPC (internal).
- **Functionality**: Accepts HTTP requests, forwards them to the internal network (Processor), and serves the Swagger UI for easy API exploration and testing.

### 2. Processor Service

- **Role**: Business logic hub and orchestration layer.
- **Transports**: gRPC (Server & Client).
- **Functionality**: Currently acts as an intermediary, but designed to eventually accumulate, aggregate, and cache repository data before delivering it to the Gateway.

### 3. Collector Service

- **Role**: Data acquisition from external sources.
- **Transports**: gRPC (Server) and REST (GitHub API Client).
- **Functionality**: Encapsulates all GitHub API interaction logic. It takes a repository target (owner/repo) and fetches the raw statistics.

### 4. Subscriber Service

- **Role**: Example background worker.

## Endpoints

### Health Check & Service Discovery

`GET /api/ping`

Monitors the availability of the entire internal chain (Processor and Subscriber).

- Returns `200 OK` with statuses if all services are reachable.
- Returns `503 Service Unavailable` if the system is degraded.

### Repository Statistics

`GET /api/repositories/info?url=<github_repo_url>`

Fetches comprehensive data about a specific GitHub repository, including its name, description, star/fork counts, and creation date.

**Example Request:**

```http
GET /api/repositories/info?url=https://github.com/golang/go
```

## Getting Started

### Quick Start

The recommended way to run the entire system:

```shell
docker compose up --build
```

The API Gateway will be available at `http://localhost:28080`.

### Manual Protobuf Generation

If you modify `.proto` definitions, regenerate the gRPC code:

```shell
make protobuf
```

## Documentation

Explore the API and its data structures via Swagger UI:

```md
http://localhost:28080/swagger/index.html
```

> Don't forget to run ```cd api && swag init -g cmd/app/main.go --parseDependency --parseInternal``` inside the root directory if you update API comments.

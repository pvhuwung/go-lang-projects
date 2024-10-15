# Go Load Balancer

## Overview

This project implements a simple load balancer in Go using Domain-Driven Design (DDD) principles. It efficiently distributes incoming traffic among multiple backend servers. (On testing)

## Project Structure

```
 /load-balancer
├── cmd
│   └── main.go         # Application entry point
├── domain
│   ├── traffic         # Handles traffic distribution
│   │   ├── model.go    # Defines request and server models
│   │   ├── handler.go  # Logic for handling incoming traffic
│   └── backend         # Backend management (servers, health checks)
│       ├── model.go    # Backend server models
│       ├── manager.go  # Backend server manager
│       └── health.go   # Health check manager
├── infra               # Infrastructure concerns (HTTP, networking)
│   ├── http.go         # HTTP server setup for the load balancer
└── README.md
```

## Getting Started

Run the load balancer:
   ```bash
   cd load-balancer
   go run cmd/main.go
   ```
  

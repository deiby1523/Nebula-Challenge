# TLS Security Checker (Go)
A high-performance command-line tool written in Go for analyzing TLS/SSL security configurations across multiple domains with concurrent execution.

## Features
- Concurrent Domain Analysis: Process multiple domains simultaneously using goroutines

- Real-time Results: Instant feedback as each domain is processed

- Performance Metrics: Measure and display response times per domain

- Clean Console Output: Well-formatted, readable results

- Extensible Architecture: Designed for easy integration with external APIs (SSL Labs, etc.)

- Idiomatic Go: Focus on standard libraries and best practices

## Prerequisites
- Go 1.20 or higher

- Internet connection

- Git (for installation)

## Installation & Setup

### Clone the repository
```bash
git clone https://github.com/your-username/tls-security-checker-go.git
cd tls-security-checker-go
```

### Initialize and download dependencies
```bash
go mod tidy
```

### Run directly
```bash
go run main.go
```

## Architecture

### Concurrency Model
- Fan-out/Fan-in Pattern: Each domain processed in its own goroutine

- Channel-based Communication: Results transmitted via buffered channels

- Synchronized Collection: Main routine aggregates results efficiently

- Controlled Parallelism: Configurable worker pool for optimal performance

### Performance Characteristics
- Memory Efficient: Minimal allocations per goroutine

- Scalable: Handles hundreds of domains concurrently

- Responsive: Real-time progress feedback

- Robust: Graceful error handling and cleanup

👤 Author
Deiby Prada

📄 License
This project is licensed under the MIT License - see the LICENSE file for details.

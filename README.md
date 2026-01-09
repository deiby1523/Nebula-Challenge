# TLS Security Checker
A high-performance command-line tool written in Go for analyzing TLS/SSL security configurations in a given domain.

## Features

- Clean Console Output: Well-formatted, readable results

- Idiomatic Go: Focus on standard libraries and best practices

## Prerequisites
- Go 1.20 or higher

- Internet connection

- Git (for installation)

## Installation & Setup

### Clone the repository
```bash
git clone https://github.com/deiby1523/Nebula-Challenge.git
cd Nebula-Challenge
```

### Initialize and download dependencies
```bash
go mod tidy
```

### build the app
```bash
go build cmd/tls-checker/main.go
```

This will generate an executable file called main.exe. You can run this file with the next command

```bash
.\main.exe
```

### Run directly
If you don't want to generate an executable .exe file, you can also run the application directly from the terminal with the following command
```bash
go run cmd/tls-checker/main.go
```
### or run with an argument
You can also pass an argument to this command if you prefer.
```bash
go run cmd/tls-checker/main.go <domain>
```
### Example
```bash
go run cmd/tls-checker/main.go www.uts.edu.co
```

👤 Author
Deiby Prada

📄 License
This project is licensed under the MIT License - see the LICENSE file for details.

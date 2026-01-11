# TLS Security Checker

A command-line tool written in Go to analyze the TLS/SSL configuration of a domain using the SSL Labs public API.

## Features

- Clean console output: results formatted for human reading
- CLI with subcommands: `analyze` and `list`
- Idiomatic in Go: use of standard libraries and modular structure

## Requeriments

- Go 1.20 or higher
- Internet connection
- Git (optional, to clone the repository)

## Installation and usage
Clone the repository:

```bash
git clone https://github.com/deiby1523/Nebula-Challenge.git
cd Nebula-Challenge
```

Download dependencies:

```bash
go mod tidy
```

Build the app:

```bash
go build -o tls-checker.exe ./cmd/tls-checker
```

Generated executables: `tls-checker.exe` (Windows)

### App info

To display an informational message about the application and how to use it correctly, you can use the following command. 

If you are in the Visual Studio Code (PowerShell) terminal, you need to put a dot './' before the command, but if you are in an external terminal like cmd, you don't need to add a dot './' at the beginning.

```bash
./tls-checker
```


### Using the commands

- Analyze a domain by passing the domain through a flag:

```bash

./tls-checker analyze --domain www.example.com

# or also using go run
go run ./cmd/tls-checker analyze --domain www.example.com
```

- Analyze a domain by requesting the domain via terminal (without flag):

```bash
./tls-checker analyze
# o
go run ./cmd/tls-checker analyze
```

The domain will be requested via stdin if it is not provided with `--domain`.

- List saved results (reads results.json and displays the records):

```bash
./tls-checker list
# o
go run ./cmd/tls-checker list
```

### Saving results

After completing an analysis, the program asks if you want to save the result. If you confirm, the result is added to the `results.json` file in the project root.


## Testing
To run unit tests for the project:

```bash
go test ./...
```

## Autor

Deiby Prada


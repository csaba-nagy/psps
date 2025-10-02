# psps 🐈

A simple and efficient port scanning tool written in Go that allows you to scan TCP ports on a target host.

## Features

- Scan TCP ports on any host
- Configurable port range
- Parallel scanning using multiple workers
- Graceful shutdown handling
- Flexible output options (console or file)
- Simple and intuitive command-line interface

## Installation

To install the port scanner, you can either build it from source or download a pre-built binary.

From source:
```bash
go install github.com/csaba-nagy/psps/cmd/psps@latest
```

## Usage

Basic usage:
```bash
psps [options]
```

### Command Options

The following command-line options are available:

- `-host`: Target host to scan (default: "127.0.0.1")
- `-from`: Starting port number (default: 8080)
- `-to`: Ending port number (default: 8090)
- `-workers`: Number of concurrent workers (default: number of CPU cores)
- `-output`: Path to output file (optional). If specified, results will be written to this file instead of console output.

Example usage:
```bash
# Scan ports 8080-8090 on localhost (default)
psps

# Scan ports 20-100 on example.com with 4 workers
psps -host example.com -from 20 -to 100 -workers 4

# Scan ports 1000-2000 on a specific IP and save results to a file
psps -host 192.168.1.1 -from 1000 -to 2000 -output results.txt
```

## License

MIT License

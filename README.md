# Remora

**Remora** is a high-performance, Redis-like in-memory key-value store written in Go. It features a concurrent server architecture and communicates using the Remora Serialization Protocol (RESP). It comes with a dedicated interactive command-line client for easy interaction.

## ✨ Features

- **In-Memory Storage**: Fast key-value store operations.
- **Redis-like Commands**: Supports common commands like `SET`, `GET`, `PING`, and `ECHO`.
- **RESP-compliant**: Uses the [Remora Serialization Protocol](./pkg/resp/README.md) for client-server communication.
- **High Concurrency**: Built to handle many concurrent client connections efficiently.
- **Graceful Shutdown**: Ensures data integrity and clean exits.
- **Interactive CLI**: Includes a `remora-cli` for a user-friendly REPL experience.

## 🚀 Getting Started

### Prerequisites

- Go (version 1.20 or later)

### Installation & Running

1.  **Clone the repository:**
    ```sh
    git clone https://github.com/arhantbararia/Remora.git
    cd Remora
    ```

2.  **Run the server:**
    The server will start on its default port.(3475)
    ```sh
    go run ./cmd/remora
    ```

3.  **Run the CLI client:**
    In a new terminal, connect to the server.
    ```sh
    go run ./cmd/remora-cli
    ```

### Example Session

Once the CLI is running, you can start sending commands:

```
> PING
PONG
> SET message "Hello, Remora!"
OK
> GET message
"Hello, Remora!"
> ECHO "Testing echo"
"Testing echo"
>
```

## 🔧 Configuration

The server can be configured via a `.env` file in the project root.

```dotenv
# Maximum number of concurrent client connections
MAX_CONCURRENT_THREADS=1000

# Timeout for reading from a client connection in minutes
READ_TIMEOUT_MINUTES=2
```

## 📂 Project Structure
```
Remora
├─ .env
├─ cmd
│  ├─ remora
│  │  └─ main.go
│  └─ remora-cli
│     ├─ internal
│     │  ├─ client.go
│     │  ├─ printer.go
│     │  └─ repl.go
│     └─ main.go
├─ go.mod
├─ pkg
│  ├─ commands
│  │  ├─ echo.go
│  │  ├─ get.go
│  │  ├─ ping.go
│  │  ├─ registry.go
│  │  └─ set.go
│  ├─ resp
│  │  ├─ parser.go
│  │  ├─ parser_test.go
│  │  ├─ README.md
│  │  ├─ serializer.go
│  │  ├─ serializer_test.go
│  │  └─ value.go
│  ├─ server
│  │  └─ server.go
│  └─ store
│     └─ store.go
└─ README.md

```


## Remora Serialization Protocol

The `resp` package provides functionality for parsing and serializing data using the Remora Serialization Protocol (RESP). RESP is a simple, human-readable protocol used by Remora for communication between clients and servers. This package defines types and functions to encode and decode RESP messages, supporting all standard RESP data types: Simple Strings, Errors, Integers, Bulk Strings, and Arrays.
[Read More](./pkg/resp/README.md)
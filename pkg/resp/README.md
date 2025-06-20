# Remora Serialization Protocol

The `resp` package provides functionality for parsing and serializing data using the Remora Serialization Protocol (RESP). RESP is a simple, human-readable protocol used by Remora for communication between clients and servers. This package defines types and functions to encode and decode RESP messages, supporting all standard RESP data types: Simple Strings, Errors, Integers, Bulk Strings, and Arrays.

The package is structured to allow easy integration into Redis-like servers or clients, enabling robust and correct handling of RESP messages, including error checking and protocol compliance.

---

## Table of Contents

- [Parser Functions](#parser-functions)
- [Serializer Functions](#serializer-functions)

---

## Parser Functions

These functions handle reading and parsing RESP messages from a buffered reader.

### `ParseRESP(r *bufio.Reader) (Value, error)`

Parses a single RESP value from the provided buffered reader. Determines the RESP type from the first byte and delegates to the appropriate parsing function. Supports Simple Strings (`+`), Errors (`-`), Integers (`:`), Bulk Strings (`$`), and Arrays (`*`). Returns a `Value` or an error if the input is malformed or incomplete.

**Example:**
```go
input := "+OK\r\n"
value, err := ParseRESP(bufio.NewReader(strings.NewReader(input)))
// value: SimpleString("OK"), err: nil
```

---

### `parseSimpleString(r *bufio.Reader) (Value, error)`

Parses a RESP Simple String ending with `\r\n` and containing no control characters. Returns a `SimpleString` value or an error if invalid.

**Example:**
```go
input := "+PONG\r\n"
value, err := parseSimpleString(bufio.NewReader(strings.NewReader(input)))
// value: SimpleString("PONG"), err: nil
```

---

### `parseError(r *bufio.Reader) (Value, error)`

Parses a RESP Error ending with `\r\n` and containing no control characters. Returns an `ErrorType` value or an error if invalid.

**Example:**
```go
input := "-ERR unknown command\r\n"
value, err := parseError(bufio.NewReader(strings.NewReader(input)))
// value: ErrorType("ERR unknown command"), err: nil
```

---

### `parseInteger(r *bufio.Reader) (Value, error)`

Parses a RESP Integer followed by `\r\n`. Returns an `Integer` value or an error if the format is invalid or not a valid integer.

**Example:**
```go
input := ":1000\r\n"
value, err := parseInteger(bufio.NewReader(strings.NewReader(input)))
// value: Integer(1000), err: nil
```

---

### `parseBulkString(r *bufio.Reader) (Value, error)`

Parses a RESP Bulk String by reading the length prefix and string data, ensuring correct length and line endings. Handles null bulk strings (length `-1`). Returns a `BulkString` value or an error if invalid.

**Example:**
```go
input := "$6\r\nfoobar\r\n"
value, err := parseBulkString(bufio.NewReader(strings.NewReader(input)))
// value: BulkString("foobar"), err: nil
```

---

### `parseArray(r *bufio.Reader) (Value, error)`

Parses a RESP Array by reading the array length and recursively parsing each element. Handles null arrays (length `-1`) and empty arrays. Returns an `Array` value or an error if invalid.

**Example:**
```go
input := "*2\r\n$3\r\nfoo\r\n$3\r\nbar\r\n"
value, err := parseArray(bufio.NewReader(strings.NewReader(input)))
// value: Array{BulkString("foo"), BulkString("bar")}, err: nil
```

---

## Serializer Functions

These functions handle serializing and writing RESP messages to a buffered writer.

### `WriteRESP(w *bufio.Writer, value Value) error`

Serializes a `Value` as a RESP message and writes it to the buffered writer. Dispatches to the appropriate serialization function based on the value's type. Returns an error if the value type is unknown or writing fails.

**Example:**
```go
var buf bytes.Buffer
err := WriteRESP(bufio.NewWriter(&buf), SimpleString("OK"))
// buf.String(): "+OK\r\n", err: nil
```

---

### `WriteSimpleString(writer *bufio.Writer, value Value) error`

Serializes a `SimpleString` value as a RESP Simple String and writes it to the writer. Ensures the string contains no control characters. Returns an error if invalid or writing fails.

**Example:**
```go
var buf bytes.Buffer
err := WriteSimpleString(bufio.NewWriter(&buf), SimpleString("PONG"))
// buf.String(): "+PONG\r\n", err: nil
```

---

### `WriteError(writer *bufio.Writer, value Value) error`

Serializes an `ErrorType` value as a RESP Error and writes it to the writer. Ensures the error string contains no control characters. Returns an error if invalid or writing fails.

**Example:**
```go
var buf bytes.Buffer
err := WriteError(bufio.NewWriter(&buf), ErrorType("ERR unknown command"))
// buf.String(): "-ERR unknown command\r\n", err: nil
```

---

### `WriteInteger(writer *bufio.Writer, value Value) error`

Serializes an `Integer` value as a RESP Integer and writes it to the writer. Returns an error if writing fails.

**Example:**
```go
var buf bytes.Buffer
err := WriteInteger(bufio.NewWriter(&buf), Integer(1000))
// buf.String(): ":1000\r\n", err: nil
```

---

### `WriteBulkString(writer *bufio.Writer, value Value) error`

Serializes a `BulkString` value as a RESP Bulk String and writes it to the writer. Handles null bulk strings and ensures correct length and line endings. Returns an error if writing fails.

**Example:**
```go
var buf bytes.Buffer
err := WriteBulkString(bufio.NewWriter(&buf), BulkString("foobar"))
// buf.String(): "$6\r\nfoobar\r\n", err: nil
```

---

### `WriteArray(writer *bufio.Writer, value Value) error`

Serializes an `Array` value as a RESP Array and writes it to the writer. Handles null and empty arrays, and recursively serializes each element. Returns an error if writing fails.

**Example:**
```go
var buf bytes.Buffer
arr := Array{BulkString("foo"), BulkString("bar")}
err := WriteArray(bufio.NewWriter(&buf), arr)
// buf.String(): "*2\r\n$3\r\nfoo\r\n$3\r\nbar\r\n", err: nil
```

---
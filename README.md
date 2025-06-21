# Remora

## RESP

The Remora Serialization Protocol (RESP) is a simple, human-readable protocol used by Remora for communication between clients and servers. RESP supports several data types, including Simple Strings, Errors, Integers, Bulk Strings, and Arrays, enabling robust and correct message handling for Redis-like systems.

[the RESP package documentation](./pkg/resp/README.md).
```
Remora
├─ cmd
│  └─ remora
│     └─ main.go
├─ go.mod
├─ pkg
│  └─ resp
│     ├─ parser.go
│     ├─ parser_test.go
│     ├─ README.md
│     ├─ serializer.go
│     ├─ serializer_test.go
│     └─ value.go
└─ README.md

```
package internal

import (
	"bufio"
	"net"

	"remora/pkg/resp"
)

type Connection struct {
	Reader *bufio.Reader
	Writer *bufio.Writer
}

func wrapConnection(conn *net.Conn, connection *Connection) {
	if conn == nil || connection == nil {
		return
	}
	connection.Reader = bufio.NewReader(*conn)
	connection.Writer = bufio.NewWriter(*conn)
}

func Connect(host, port string) (Connection, error) {
	connection := Connection{}

	address := net.JoinHostPort(host, port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return connection, err
	}
	wrapConnection(&conn, &connection)

	return connection, nil

}

func sendCommand(conn Connection, cmd string, args []string) error {

	elems := append([]string{cmd}, args...)
	values := make([]resp.Value, len(elems))
	for i, s := range elems {
		values[i] = resp.Value{Type: resp.BulkString, Bulk: []byte(s)}
	}
	arrayVal := resp.Value{Type: resp.Array, Array: values}

	// Write RESP Array to conn.Writer
	if err := resp.WriteRESP(conn.Writer, arrayVal); err != nil {
		return err
	}

	// Flush the writer
	return conn.Writer.Flush()
}

func recieveResponse(conn Connection) (resp.Value, error) {
	value, err := resp.ParseRESP(conn.Reader)
	if err != nil {
		return resp.Value{}, err
	}
	return value, nil
}

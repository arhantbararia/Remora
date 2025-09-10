package server

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"remora/pkg/commands"
	"remora/pkg/resp"
	"remora/pkg/store"
	"strings"
	"time"
)

const (
	maxConcurrentConnections = 1000
	connectionReadTimeout    = 2 * time.Minute
)

type Server struct {
	Host      string
	Port      string
	store     *store.Store
	semaphore chan struct{}
}

func (s *Server) ListenAndServe() error {

	address := fmt.Sprintf("%v:%v", s.Host, s.Port)

	//initialize the Listener
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", address, err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}
		log.Printf("Client connected: %s", conn.RemoteAddr().String())

		go s.handleConnection(conn)

	}

}

func NewRemoraServer(host string, port string) *Server {

	return &Server{
		Host:      host,
		Port:      port,
		store:     store.NewStore(),
		semaphore: make(chan struct{}, maxConcurrentConnections),
	}
}

func connectionClose(conn net.Conn) {
	log.Printf("Client disconnected: %s", conn.RemoteAddr().String())
	conn.Close()
}

func (s *Server) handleConnection(conn net.Conn) {

	defer connectionClose(conn)

	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v", r)
		}
		connectionClose(conn)
	}()

	reader := bufio.NewReader(conn) //ex: ECHO "HELLO WORLD"
	writer := bufio.NewWriter(conn)
	// store := store.NewStore()
	for {

		conn.SetReadDeadline(time.Now().Add(connectionReadTimeout))

		// 1. Parse the next RESP message
		value, err := resp.ParseRESP(reader) //ex: Value{Type: Array , Array: ["$ECHO", "$"HELLO WORLD""]}
		if err != nil {
			if err == io.EOF {
				//client closed the connection
				continue
			}

			resp.WriteError(writer, resp.Value{
				Type: resp.ErrorType,
				Str:  "ERR Protocol error: " + err.Error(),
			})
			writer.Flush()
			return
		}

		if value.Type != resp.Array { //check if value.Type is Array since commands and parametrs are space sperated array items
			resp.WriteError(writer, resp.Value{
				Type: resp.ErrorType,
				Str:  "ERR Protocol error: expected Array",
			})
			writer.Flush()
			return
		}

		// Extract command name

		//check if first element is bulk string
		if len(value.Array) == 0 || value.Array[0].Type != resp.BulkString {
			resp.WriteError(writer, resp.Value{
				Type: resp.ErrorType,
				Str:  "ERR Protocol Error: invalid format",
			})
			writer.Flush()
			return
		}

		cmdName := strings.ToUpper(string(value.Array[0].Bulk)) //make command name ex. echo -> ECHO

		handler, ok := commands.GetHandler(cmdName) //get the handler function for the comman
		if !ok {
			resp.WriteError(writer, resp.Value{
				Type: resp.ErrorType,
				Str:  "ERR Protocol Error: invalid command",
			})
			writer.Flush()
			continue

		}

		reply := handler(s.store, value.Array[1:]) // pass "$HELLO WORLD" to the ECHO handler function
		//reply is of type Value

		resp.WriteRESP(writer, reply)
		writer.Flush()

	}

}

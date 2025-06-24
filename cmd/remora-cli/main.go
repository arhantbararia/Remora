package main

import (
	"flag"
	"fmt"
	"os"
	"remora/cmd/remora-cli/internal"
)

func parseFlags() (host, port *string) {
	host = flag.String("host", "", "Remora Server address (default: localhost)")
	port = flag.String("port", "", "Remora Server listening port (default: 1379)")
	help := flag.Bool("help", false, "Show help")

	flag.Parse()

	if *help {
		fmt.Println("Usage:")
		flag.PrintDefaults()
		os.Exit(0)
	}
	if *host == "" {
		*host = "localhost"
	}

	if *port == "" {
		*port = "1379"
	}
	return host, port
}

func main() {
	host, port := parseFlags()

	fmt.Printf("Connecting to server at %s:%s\n", *host, *port)

	conn, err := internal.Connect(*host, *port)
	if err != nil {
		fmt.Printf("Some error occured: %v \n", err)
		os.Exit(503)
	}
	internal.RunREPL(conn)

}

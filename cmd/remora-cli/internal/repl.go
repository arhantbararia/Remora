package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RunREPL(conn Connection) {
	for {
		fmt.Print("remora-cli> ")
		scanner := bufio.NewScanner(os.Stdin)
		if !scanner.Scan() {
			fmt.Println()
			break
		}
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		if line == "exit" || line == "quit" {
			break
		}

		tokens := strings.Fields(line)
		if len(tokens) == 0 {
			continue
		}
		cmd := tokens[0]
		args := tokens[1:]
		executeCommand(cmd, args, conn)
	}
}

func executeCommand(cmd string, args []string, conn Connection) {
	err := sendCommand(conn, cmd, args)
	if err != nil {
		fmt.Printf("Some Error Occured: %v", err)
		os.Exit(499)

	}

	value, err := recieveResponse(conn)
	if err != nil {
		fmt.Printf("Some error occurred: %v", err)
	}
	printValue(value)

}

// func parseArgs(input string) []string {
// 	var args []string
// 	var current strings.Builder
// 	inQuotes := false
// 	escaped := false

// 	for _, r := range input {
// 		switch {
// 		case escaped:
// 			current.WriteRune(r)
// 			escaped = false
// 		case r == '\\':
// 			escaped = true
// 		case r == '"':
// 			inQuotes = !inQuotes
// 		case r == ' ' && !inQuotes:
// 			if current.Len() > 0 {
// 				args = append(args, current.String())
// 				current.Reset()
// 			}
// 		default:
// 			current.WriteRune(r)
// 		}
// 	}
// 	if current.Len() > 0 {
// 		args = append(args, current.String())
// 	}
// 	return args
// }

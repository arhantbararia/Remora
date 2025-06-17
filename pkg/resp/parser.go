package parser

import (
	"bufio"
	"fmt"
)

func ParseRESP(r *bufio.Reader) (Value, error) {

	prefix, err := r.ReadByte()
	if err != nil {
		return Value{}, err
	}

	switch prefix {
	case '+':
		return parseSimpleString(r)
	
	case '-':
		return parseError(r)
	
	case ':':
		return parseInteger(r)
	
	case '$':
		return parseBulkString(r)
	case '*':
		return parseArray(r)
	default:
		return Value{}, fmt.Errorf("ERR Protocol Error: invalid_prefix '%c' ", prefix )

	
	}
}




func parseSimpleString(r *bufio.Reader) (Value , error ) {
	line, err := r.ReadString('\n')
	if err != nil {
		return Value{}, err
	}
	// line should end with \r\n
	if len(line) < 2 || line[len(line)-2] != '\r' {
		return Value{}, fmt.Errorf("ERR Protocol Error: invalid line ending in simple string")
	}
	str := line[:len(line)-2]
	// Check for control characters in the string
	for i := 0; i < len(str); i++ {
		if str[i] == '\r' || str[i] == '\n' {
			return Value{}, fmt.Errorf("ERR Protocol Error: control character in simple string")
		}
	}
	return Value{
		Type: SimpleString,
		Str:  str,
	}, nil
}
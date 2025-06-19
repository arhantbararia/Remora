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



func parseError(r *bufio.Reader ) (Value , error ) {
	line, err :=  r.ReadString('\n')
	if err != nil {
		return Value{}, err

	}

	if len(line) < 2 || line[len(line)-2] != '\r' {
		return Value{}, fmt.Errorf("ERR Protocol Error: invalid line ending in error string")

	}


	error_str := line[:len(line)-2]
	
	for i := 0; i < len(error_str); i++ {
		if error_str[i] == '\r' || error_str[i] == '\n' {
			return Value{}, fmt.Errorf("ERR Protocol Error: control character in Error string")
		}
	}


	return Value{
		Type: ErrorType,
		Str: error_str,
	}, nil

	



}




func parseInteger(r *bufio.Reader) (Value , error ) {
	line, err := r.ReadString('\n')
	if err != nil {
		return Value{}, err
	}

	if len(line) < 2 || line[len(line)-2] != '\r' {
		return Value{}, fmt.Errorf("ERR Protocol Error: invalid line ending in error string")

	}

	
	
	intStr := line[:len(line)-2]
	var n int64
	_, err = fmt.Sscanf(intStr, "%d", &n)
	if err != nil {
		return Value{}, fmt.Errorf("ERR Protocol Error: invalid integer value")
	}
	return Value{
		Type: Integer,
		Int:  n,
	}, nil

}

func parseBulkString(r *bufio.Reader) (Value, error) {
	
	//Read the length line
	lenLine, err := r.ReadString('\n')
	if err != nil {
		return Value{} , err
	}

	if len(lenLine) < 2 || lenLine[len(lenLine)-2] != '\r' {
		return Value{}, fmt.Errorf("ERR Protocaol Error: Invalid bulk string")
	}

	lenStr := lenLine[:len(lenLine)-2]
	var length int64
	_, err = fmt.Sscanf(lenStr, "%d", &length)
	
	if err != nil || length < -1 {
		return Value{}, fmt.Errorf("ERR Protocol Error: invalid length value")
	}


	if length == -1 {
		return Value{
			Type: BulkString,
			Bulk: nil,
		}, nil
	}

	//Read the string line
	strLine, err := r.ReadString('\n')
	if err != nil {
		return Value{}, err
	}

	if len(lenLine) < 2 || lenLine[len(lenLine)-2] != '\r' {
		return Value{}, fmt.Errorf("ERR Protocaol Error: Invalid bulk string")
	}

	b_string := strLine[:len(strLine)-2]

	if len(b_string) != int(length) {
		return Value{}, fmt.Errorf("ERR Protocol Error: bulk string length mismatch")
	}

	for i := 0; i < int(length); i++{
		if b_string[i] == '\r' || b_string[i] == '\n' {
			return Value{}, fmt.Errorf("ERR Protocol Error: control character in Error string")
		}
	}

	return Value{
		Type: BulkString,
		Bulk: []byte(b_string),
	}, nil




}



func parseArray(r *bufio.Reader) ( Value , error ){
	//Read the length line
	countLine, err := r.ReadString('\n')
	if err != nil {
		return Value{} , err
	}

	if len(countLine) < 2 || countLine[len(countLine)-2] != '\n' {
		return Value{}, fmt.Errorf("ERR Protocol Error: Invalid bulk string")
	}

	cntStr := countLine[:len(countLine)-2]
	var count int64
	_, err = fmt.Sscanf(cntStr, "%d", &count)
	if err != nil || count < -1 {
		return Value{}, fmt.Errorf("ERR Protocol Error: Invalid Array len count")
	}
	if count == -1 {
		return Value{
			Type: Array,
			Bulk: nil,
		}, nil
	}

	values := make([]Value, count)
	for i := int64(0); i < count; i++ {
		val, err := ParseRESP(r)
		if err != nil {
			return Value{}, err
		}
		values[i] = val
	}
	return Value{
		Type:  Array,
		Array: values,
	}, nil

}




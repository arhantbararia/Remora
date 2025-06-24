package commands

import "remora/pkg/resp"

func echoHandler(args []resp.Value) resp.Value {

	if len(args) == 0 {
		return resp.Value{
			Type: resp.ErrorType,
			Str:  "ERR: wrong number of arguments for ECHO",
		}
	}

	// Concatenate all Bulk strings from args
	var totalLen int
	for _, arg := range args {
		if arg.Type != resp.BulkString || arg.Bulk == nil {
			return resp.Value{
				Type: resp.ErrorType,
				Str:  "ERR: arguments must be bulk strings",
			}
		}
		totalLen += len(arg.Bulk)
	}
	joined := make([]byte, 0, totalLen+len(args)-1)
	for i, arg := range args {
		if i > 0 {
			joined = append(joined, ' ')
		}
		joined = append(joined, arg.Bulk...)
	}

	return resp.Value{
		Type: resp.BulkString,
		Bulk: joined,
	}

}

func init() {
	Register("ECHO", echoHandler)
}

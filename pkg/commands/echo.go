package commands

import "remora/pkg/resp"

func echoHandler(args []resp.Value ) resp.Value {

	if len(args) != 1 {
		return resp.Value{
			Type: resp.ErrorType,
			Str: "ERR: wrong number of arguments for ECHO",
		}
	}

	if args[0].Type != resp.BulkString {
		return resp.Value{
			Type: resp.ErrorType,
			Str: "ERR: arguments must be a bulk string",
		}
	}

	return resp.Value{
		Type: resp.BulkString,
		Bulk: args[0].Bulk,
	}

}

 
func init(){
	Register("ECHO" , echoHandler)
}
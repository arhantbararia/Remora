package commands

import "remora/pkg/resp"

func pingHandler(args []resp.Value ) resp.Value {

	switch(len(args)) {
	case 0:
		return resp.Value{
			Type: resp.SimpleString,
			Str: "PONG",
		}
	
	case 1: //behave like echo
		if args[0].Type != resp.BulkString {
			return resp.Value{
				Type: resp.ErrorType,
				Str: "ERR: argument must be bulk string",
			}
		}
		return resp.Value{
			Type: resp.BulkString,
			Bulk: args[0].Bulk,
		}
	default:
		return resp.Value{
			Type: resp.ErrorType,
			Str: "ERR: Wrong number of arguments for 'PING' ",
		}

	}
}

func init() {
	Register("PING", pingHandler)
}



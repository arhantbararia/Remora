package commands

import (
	"remora/pkg/resp"
	"remora/pkg/store"
)

func getHandler(store *store.Store, args []resp.Value) resp.Value {
	if store == nil {
		return resp.Value{
			Type: resp.ErrorType,
			Str:  "ERR: store is not initialized",
		}
	}

	if len(args) != 1 {
		return resp.Value{
			Type: resp.ErrorType,
			Str:  "ERR wrong number of arguments for 'GET' command",
		}
	}

	keyArg := args[0]
	if keyArg.Type != resp.BulkString {
		return resp.Value{
			Type: resp.ErrorType,
			Str:  "ERR wrong argument type for 'GET' command",
		}
	}

	value, ok := store.Get(keyArg.Str)
	if !ok {
		return resp.Value{
			Type: resp.BulkString,
			Bulk: nil, // Return nil for non-existent keys
		}
	}
	return value

}

func init() {
	Register("GET", getHandler) // GET uses store
}

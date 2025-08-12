package commands

import (
	"remora/pkg/resp"
	"remora/pkg/store"
)

func setHandler(store *store.Store, args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{
			Type: resp.ErrorType,
			Str:  "ERR wrong number of arguments for 'SET' command",
		}
	}

	keyArg, value := args[0], args[1]
	if keyArg.Type != resp.BulkString || !(value.Type == resp.BulkString || value.Type == resp.Integer || value.Type == resp.SimpleString) {
		return resp.Value{
			Type: resp.ErrorType,
			Str:  "ERR wrong argument type for 'SET' command",
		}
	}

	key := keyArg.Str

	store.Set(key, value)

	return resp.Value{
		Type: resp.SimpleString,
		Str:  "OK",
	}

}

func init() {
	Register("SET", setHandler) // SET uses store
}

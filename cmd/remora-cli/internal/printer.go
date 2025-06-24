package internal

import "remora/pkg/resp"

func printValue(value resp.Value) {
	switch value.Type {
	case resp.SimpleString:
		println(value.Str)
	case resp.BulkString:
		if value.Bulk == nil {
			println("<nil>")
		} else {
			println(string(value.Bulk))
		}
	case resp.Integer:
		println(value.Int)
	case resp.ErrorType:
		println("ERR:", value.Str)
	case resp.Array:
		for _, elem := range value.Array {
			printValue(elem)
		}
	}
}

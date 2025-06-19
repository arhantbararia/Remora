package parser

type RespType int

const (
	SimpleString RespType = iota
	ErrorType
	Integer
	BulkString
	Array
)

type Value struct {
	Type  RespType
	Str   string
	Int   int64
	Bulk  []byte  // for bulk string (nil if null)
	Array []Value // for array (nil if null array, or empty slice if *0)
}

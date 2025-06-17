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
	Type RespType
	Str string
	Int int64
	Bulk []byte
	Array []Value
}
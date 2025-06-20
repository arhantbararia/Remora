package resp

import (
	"bufio"
	"bytes"
	"testing"
)

func TestWriteRESP(t *testing.T) {
	tests := []struct {
		name    string
		value   Value
		want    string
		wantErr bool
	}{
		{
			name:  "SimpleString",
			value: Value{Type: SimpleString, Str: "OK"},
			want:  "+OK\r\n",
		},
		{
			name:    "SimpleStringWithNewline",
			value:   Value{Type: SimpleString, Str: "foo\nbar"},
			wantErr: true,
		},
		{
			name:  "Error",
			value: Value{Type: ErrorType, Str: "ERR something went wrong"},
			want:  "-ERR something went wrong\r\n",
		},
		{
			name:    "ErrorWithCR",
			value:   Value{Type: ErrorType, Str: "foo\rbar"},
			wantErr: true,
		},
		{
			name:  "IntegerPositive",
			value: Value{Type: Integer, Int: 123},
			want:  ":123\r\n",
		},
		{
			name:  "IntegerNegative",
			value: Value{Type: Integer, Int: -42},
			want:  ":-42\r\n",
		},
		{
			name:  "BulkStringNormal",
			value: Value{Type: BulkString, Bulk: []byte("foobar")},
			want:  "$6\r\nfoobar\r\n",
		},
		{
			name:  "BulkStringEmpty",
			value: Value{Type: BulkString, Bulk: []byte("")},
			want:  "$0\r\n\r\n",
		},
		{
			name:  "BulkStringNull",
			value: Value{Type: BulkString, Bulk: nil},
			want:  "$-1\r\n",
		},
		{
			name: "ArrayOfSimpleStrings",
			value: Value{
				Type: Array,
				Array: []Value{
					{Type: SimpleString, Str: "foo"},
					{Type: SimpleString, Str: "bar"},
				},
			},
			want: "*2\r\n+foo\r\n+bar\r\n",
		},
		{
			name:  "NullArray",
			value: Value{Type: Array, Array: nil},
			want:  "*-1\r\n",
		},
		{
			name:  "EmptyArray",
			value: Value{Type: Array, Array: []Value{}},
			want:  "*0\r\n",
		},
		{
			name:    "UnknownType",
			value:   Value{Type: 99},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			w := bufio.NewWriter(&buf)
			err := WriteRESP(w, tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteRESP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && buf.String() != tt.want {
				t.Errorf("WriteRESP() = %q, want %q", buf.String(), tt.want)
			}
		})
	}
}

package resp

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestParseRESP(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Value
		wantErr bool
	}{
		{
			name:  "SimpleString",
			input: "+OK\r\n",
			want:  Value{Type: SimpleString, Str: "OK"},
		},
		{
			name:  "Error",
			input: "-ERR something went wrong\r\n",
			want:  Value{Type: ErrorType, Str: "ERR something went wrong"},
		},
		{
			name:  "Integer",
			input: ":1000\r\n",
			want:  Value{Type: Integer, Int: 1000},
		},
		{
			name:  "BulkString",
			input: "$6\r\nfoobar\r\n",
			want:  Value{Type: BulkString, Bulk: []byte("foobar")},
		},
		{
			name:  "NullBulkString",
			input: "$-1\r\n",
			want:  Value{Type: BulkString, Bulk: nil},
		},
		{
			name:  "EmptyBulkString",
			input: "$0\r\n\r\n",
			want:  Value{Type: BulkString, Bulk: []byte("")},
		},
		{
			name:  "Array of SimpleStrings",
			input: "*2\r\n+foo\r\n+bar\r\n",
			want: Value{
				Type: Array,
				Array: []Value{
					{Type: SimpleString, Str: "foo"},
					{Type: SimpleString, Str: "bar"},
				},
			},
		},
		{
			name:  "NullArray",
			input: "*-1\r\n",
			want:  Value{Type: Array, Array: nil},
		},
		{
			name:  "EmptyArray",
			input: "*0\r\n",
			want:  Value{Type: Array, Array: []Value{}},
		},
		{
			name:    "InvalidPrefix",
			input:   "#foo\r\n",
			wantErr: true,
		},
		{
			name:    "InvalidSimpleStringEnding",
			input:   "+foo\n",
			wantErr: true,
		},
		{
			name:    "InvalidBulkStringLength",
			input:   "$abc\r\nfoobar\r\n",
			wantErr: true,
		},
		{
			name:    "BulkStringLengthMismatch",
			input:   "$3\r\nfoobar\r\n",
			wantErr: true,
		},
		{
			name:    "InvalidArrayLength",
			input:   "*abc\r\n",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := bufio.NewReader(strings.NewReader(tt.input))
			got, err := ParseRESP(r)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"\nTest: %s\nInput: %q\nExpected error: %v\nActual error: %v\nExpected: %#v\nActual: %#v",
					tt.name, tt.input, tt.wantErr, err, tt.want, got,
				)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf(
					"\nTest: %s\nInput: %q\nExpected: %#v\nActual: %#v",
					tt.name, tt.input, tt.want, got,
				)
			}
			if tt.wantErr && err != nil {
				t.Logf(
					"PASS: %s (expected error and got error: %v)\nInput: %q",
					tt.name, err, tt.input,
				)
			}
			if !tt.wantErr && reflect.DeepEqual(got, tt.want) {
				t.Logf(
					"PASS: %s\nInput: %q\nOutput: %#v",
					tt.name, tt.input, got,
				)
			}
		})
	}
}

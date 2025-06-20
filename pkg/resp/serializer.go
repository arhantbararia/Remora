package resp

import (
	"bufio"
	"io"
	"strconv"
)

func WriteRESP(w *bufio.Writer, value Value) error {
	switch value.Type {
	case SimpleString:
		return WriteSimpleString(w, value)
	case BulkString:
		return WriteBulkString(w, value)
	case ErrorType:
		return WriteError(w, value)
	case Integer:
		return WriteInteger(w, value)
	case Array:
		return WriteArray(w, value)
	default:
		return io.ErrUnexpectedEOF
	}
}

func WriteSimpleString(writer *bufio.Writer, value Value) error {

	str := value.Str
	for i := 0; i < len(str); i++ {
		if str[i] == '\r' || str[i] == '\n' {
			return io.ErrUnexpectedEOF
		}
	}
	if _, err := writer.WriteString("+" + str + "\r\n"); err != nil {
		return err
	}
	return writer.Flush()

}

func WriteError(writer *bufio.Writer, value Value) error {
	errorStr := value.Str

	for i := 0; i < len(errorStr); i++ {
		if errorStr[i] == '\r' || errorStr[i] == '\n' {
			return io.ErrUnexpectedEOF
		}
	}

	if _, err := writer.WriteString("-" + errorStr + "\r\n"); err != nil {
		return err
	}
	return writer.Flush()

}

func WriteInteger(writer *bufio.Writer, value Value) error {
	intVal := value.Int

	if _, err := writer.WriteString(":" + strconv.FormatInt(intVal, 10) + "\r\n"); err != nil {
		return err
	}

	return writer.Flush()

}

func WriteBulkString(writer *bufio.Writer, value Value) error {
	if value.Bulk == nil {
		if _, err := writer.WriteString("$-1\r\n"); err != nil {
			return err
		}
		return writer.Flush()
	}

	data := value.Bulk
	length := len(value.Bulk)
	if _, err := writer.WriteString("$" + strconv.Itoa(length) + "\r\n"); err != nil {
		return err
	}
	if _, err := writer.Write(data); err != nil { //directly write bytes no string conversion required
		return err
	}
	if _, err := writer.WriteString("\r\n"); err != nil {
		return err
	}
	return writer.Flush()
}

func WriteArray(writer *bufio.Writer, value Value) error {
	if value.Array == nil {
		if _, err := writer.WriteString("*-1\r\n"); err != nil {
			return err
		}

		return writer.Flush()
	}

	length := len(value.Array)
	if _, err := writer.WriteString("*" + strconv.Itoa(length) + "\r\n"); err != nil {
		return err
	}

	for _, item := range value.Array {
		if err := WriteRESP(writer, item); err != nil {
			return err
		}
	}

	return writer.Flush()

}

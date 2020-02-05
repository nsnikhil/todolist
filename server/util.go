package server

import (
	"bytes"
	"encoding/binary"
	"todolist/applogger"
)

func toByteSlice(n int64) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, n)
	if err != nil {
		applogger.Errorf("%s : %v", "[server] [toByteSlice]", err)
		return []byte{}
	}
	return buf.Bytes()
}

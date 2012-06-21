package vm

import (
	//"io"
	"bytes"
	"encoding/binary"
)

type codeReader struct {
	buffer *bytes.Reader
}

func NewCodeReader(data []byte) *codeReader {
	cr := new(codeReader)
	cr.buffer = bytes.NewReader(data)
	return cr
}

func (cr *codeReader) setPos(pos int64) {
	cr.buffer.Seek(pos, 0)
}

func (cr *codeReader) readString(size uint32) (*string, error) {
	var b []byte = make([]byte, size)
	_, err := cr.buffer.Read(b)
	if err != nil {
		return nil, err
	}
	str := string(b)
	return &str, nil
}

func (cr *codeReader) readByte() (byte, error) {
	var res byte
	err := binary.Read(cr.buffer, binary.LittleEndian, &res)
	return res, err
}

func (cr *codeReader) readWord() (uint16, error) {
	var res uint16
	err := binary.Read(cr.buffer, binary.LittleEndian, &res)
	return res, err
}

func (cr *codeReader) readDWord() (uint32, error) {
	var res uint32
	err := binary.Read(cr.buffer, binary.LittleEndian, &res)
	return res, err
}

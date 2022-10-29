package utils

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

func Encode(message string) []byte {
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)
	binary.Write(pkg, binary.LittleEndian, length)
	binary.Write(pkg, binary.LittleEndian, []byte(message))
	return pkg.Bytes()
}

func Decode(reader *bufio.Reader) (string, error) {
	lengthByte, _ := reader.Peek(4)
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}

	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}
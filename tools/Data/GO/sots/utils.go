package sots

import (
	"encoding/binary"
	"io"
)

func Fread(f io.Reader, num uint64) ([]byte, uint64) {
	if num == 0 {
		return make([]byte, 0), 0
	}
	
	var tmp = make([]byte, num)
	var _num,_ = f.Read(tmp)
	return tmp, uint64(_num)
}

func FreadLU32(f io.Reader) (uint32) {
	var tmp = make([]byte, 4)
	f.Read(tmp)
	return binary.LittleEndian.Uint32(tmp)
}

func FreadU8(f io.Reader) (uint8) {
	var tmp = make([]byte, 1)
	f.Read(tmp)
	return uint8(tmp[0])
}

func fastGetBtLU16(val int) []byte {
	return []byte { byte(val & 0xFF), byte((val >> 8)  & 0xFF) }
}

func fastGetBtLU32(val uint32) []byte {
	return []byte { byte(val & 0xFF), byte((val >> 8)  & 0xFF), byte((val >> 16)  & 0xFF),byte((val >> 24)  & 0xFF), }
}
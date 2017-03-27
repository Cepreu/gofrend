package utils

import (
	"hash/crc32"
	"testing"
)

var castagnoliTable = crc32.MakeTable(crc32.Castagnoli)

func TestHash(t *testing.T) {
	var v int
	v = HashCode("foobar")
	if v != 3026088333 {
		t.Error("Expected 3026088333, got ", v)
	}

	crc := crc32.New(castagnoliTable)
	crc.Write([]byte("abcd"))
	v2 := crc.Sum32() 
	if v2 != 0x92c80a31 {
		t.Error("Expected 92c80a31, got ", v2)
	}
}

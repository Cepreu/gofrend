package utils

import (
	"hash/crc32"
	"testing"
)

var castagnoliTable = crc32.MakeTable(crc32.Castagnoli)

func TestHash(t *testing.T) {
	var v uint64
	v = HashCode("foobar")
	if v != uint64(3026088333) {
		t.Error("Expected 3026088333, got ", v)
	}

	crc := crc32.New(castagnoliTable)
	crc.Write([]byte("abcd"))
	v2 := crc.Sum32()
	if v2 != 0x92c80a31 {
		t.Error("Expected 92c80a31, got ", v2)
	}
}

func TestUUID(t *testing.T) {
	uid, err := GenUUIDv4()
	if err != nil {
		t.Error("Error while UUID generation")
	} else {
		t.Logf("\nTestUUID: generated id is: %s\n", uid)
	}
}

func TestDirectory(t *testing.T) {
	if CreateAgentDirectory("1111111", "myIVR2") != true {
		t.Fatal("directory was not created!!!!")
	}
}

package utils

import (
	"hash/crc32"
	"io/ioutil"
)

// HashCode calculates hash for a string in Java manner
func HashCode(s string) uint64 {
	h := uint64(0)
	for i := 0; i < len(s); i++ {
		h = 31*h + uint64(s[i])
	}
	return h
}

func getHash(filename string) (uint32, error) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	h := crc32.NewIEEE()
	h.Write(bs)
	return h.Sum32(), nil
}

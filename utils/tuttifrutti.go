package utils

import (
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"github.com/satori/go.uuid"
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

// GenUUIDv4 - returns a generated ID
func GenUUIDv4() (string, error) {
	uid, err := uuid.NewV4()
	return uid.String(), err
}

func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	} else if runtime.GOOS == "linux" {
		home := os.Getenv("XDG_CONFIG_HOME")
		if home != "" {
			return home
		}
	}
	return os.Getenv("HOME")
}

// CreateAgentDirectory - creates a directory hierarchy in the user's home
func CreateAgentDirectory(domainID string, ivrName string) bool {
	dirName := filepath.Join(userHomeDir(), "F9_"+domainID, ivrName)
	src, err := os.Stat(dirName)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(dirName, 0755)
		if errDir != nil {
			panic(err)
		}
		return true
	}

	if src.Mode().IsRegular() {
		fmt.Println(dirName, "already exist as a file!")
		return false
	}

	return false
}

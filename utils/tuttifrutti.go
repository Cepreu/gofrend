package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"unicode"

	uuid "github.com/satori/go.uuid"
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
func GenUUIDv4() string {
	uid, _ := uuid.NewV4()
	return uid.String()
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

// LogWithoutNewlines - Logs a string after replacing all newline characters
func LogWithoutNewlines(str string) {
	expression := regexp.MustCompile("\n")
	log.Print(expression.ReplaceAllString(str, " \\n "))
}

// PrettyPrint - prints golang structures
func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

// PrettyLog - logs golang structures
func PrettyLog(v interface{}) (err error) {
	b, err := json.Marshal(v)
	if err == nil {
		log.Print(string(b))
	}
	return
}

func CmdUnzip(encoded string) (string, error) {
	base64Text := make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
	base64.StdEncoding.Decode(base64Text, []byte(encoded))

	grepCmd := exec.Command("gunzip")
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write(base64Text)
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()
	return string(grepBytes), nil
}

func StripSpaces(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			// if the character is a space, drop it
			return -1
		}
		// else keep it in the string
		return r
	}, str)
}

// Strempty - returns 'true' for empty or whitespace-containing strings
func Strempty(s string) bool {
	if len(s) == 0 {
		return true
	}
	r := []rune(s)
	l := len(r)

	for l > 0 {
		l--
		if !unicode.IsSpace(r[l]) {
			return false
		}
	}
	return true
}

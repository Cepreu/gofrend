package main

import (
	"archive/zip"
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

func main() {
	encoded :=
		`H4sIAAAAAAAAAIWRwW7CMBBE73yF5XuycKhEKycIpKrqqQeg9w1ZQYSzrlgH4b9vSKUUJyB8suet
ZkZrs7jUVp3pJJXjTM/SqVbEO1dWvM/05/ormc9fXpOZVuKRS7SOKdOBRC/yiZEfwuO7pZrY5xPV
HoPen6qi8SR/Qida5P2yBf9SJzPWlLf5b1duoHvGE73bN9qGViikztdbpomT7VrDTQjEKQaGVUzl
qb6tJRiWEtUfBwsMyMCk1z1d/F2v556x9yNYuDLkmwPyUQXXpOrDuVIVgVIDHRpXgoedDAx3AeNl
9EMtjD76F/G+H+k1AgAA`
	base64Text := make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
	base64.StdEncoding.Decode(base64Text, []byte(encoded))

	//	encodedHex := hex.EncodeToString(base64Text)

	stdoutDumper := hex.Dumper(os.Stdout)
	defer stdoutDumper.Close()
	stdoutDumper.Write(base64Text)

	err := ioutil.WriteFile("/tmp/test.zip", base64Text, 0644)
	if err != nil {
		panic(err)
	}
	Cmd(base64Text)
	ExampleReader()

	b := bytes.NewReader(base64Text)
	r, err := zlib.NewReader(b)
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, r)
	r.Close()

}
func ExampleReader() {
	// Open a zip archive for reading.
	r, err := zip.OpenReader("/tmp/test.zip")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	// Iterate through the files in the archive,
	// printing some of their contents.
	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			panic(err)
		}
		_, err = io.CopyN(os.Stdout, rc, 68)
		if err != nil {
			panic(err)
		}
		rc.Close()
		fmt.Println()
	}
}
func Cmd(text []byte) {
	cmd := exec.Command("gunzip")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin,text)
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", out)
}

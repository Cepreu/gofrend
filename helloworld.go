package main

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"

	//	"io"
	"io/ioutil"
	"os"

	"github.com/Cepreu/gofrend/xmlparser"
	"github.com/Cepreu/gofrend/web"
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

	// hexdumper(base64Text)

	// Cmd(base64Text)
	//	zlibbing(base64Text)
	xmlparser.ParseIVR(strings.NewReader("<ivrScript/>"))
	web.RunServer()
}

func zlibbing(source []byte) {
	b := bytes.NewReader(source)
	r, err := zlib.NewReader(b)
	if err != nil {
		panic(err)
	}
	enflated, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(enflated))
}

func hexdumper(source []byte) {
	stdoutDumper := hex.Dumper(os.Stdout)
	defer stdoutDumper.Close()
	stdoutDumper.Write(source)
}

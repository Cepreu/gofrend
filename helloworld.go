package main

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"encoding/hex"
	"io"
	"os"
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

	b := bytes.NewReader(base64Text)
	r, err := zlib.NewReader(b)
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, r)
	r.Close()

}

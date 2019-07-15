package fulfiller

import (
	"bufio"
	"os"
	"testing"

	"github.com/Cepreu/gofrend/ivr/xmlparser"
)

func checkNil(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}

func TestSession(t *testing.T) {
	var fname = "test_files/is_large_test.five9ivr"
	f, err := os.Open(fname)
	checkNil(err, t)
	script, err := xmlparser.NewIVRScript(bufio.NewReader(f))
	checkNil(err, t)
	session, err := loadSession("abc", script)
	checkNil(err, t)
	err = session.save()
	//checkNil(err, t)
	//err = session.delete()
	err = session.delete()
	checkNil(err, t)
}

package fulfiller

import (
	"bufio"
	"os"
	"testing"

	"github.com/Cepreu/gofrend/ivr/xmlparser"
)

func TestSession(t *testing.T) {
	var fname = "test_files/is_large_test.five9ivr"
	f, err := os.Open(fname)
	check(err)
	script, err := xmlparser.NewIVRScript(bufio.NewReader(f))
	check(err)
	session, err := loadSession("abc", script)
	check(err)
	err = session.save()
	check(err)
	err = session.delete()
	check(err)
}

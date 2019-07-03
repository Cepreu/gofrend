package preparer

import (
	"testing"
)

var filename1 = "test_files/sk_F9SOAP - Copy.five9ivr"

func TestPrepareFile(t *testing.T) {
	err := PrepareFile(filename1)
	checkNil(err, t)
}

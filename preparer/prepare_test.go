package preparer

import (
	"testing"
)

var filename1 = "test_files/is_large_test.five9ivr"

func TestPrepareFile(t *testing.T) {
	err := PrepareFile(filename1)
	checkNil(err, t)
}

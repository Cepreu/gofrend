package preparer

import (
	"testing"
)

var filename1 = "test_files/comparison_test.five9ivr"

func TestPrepareFile(t *testing.T) {
	err := Prepare("sk_text", "sergei_inbound", "sam2@007.com", "pwd1234567")
	checkNil(err, t)
}

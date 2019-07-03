package fulfiller

import (
	"io/ioutil"
	"testing"

	"github.com/Cepreu/gofrend/utils"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

var fname1 = "test_files/is_large_test.five9ivr"

func TestGetScript(t *testing.T) {
	data, err := ioutil.ReadFile(fname1)
	check(err)
	hash := utils.HashToString(data)
	_, err = getScript(hash)
	check(err)
}

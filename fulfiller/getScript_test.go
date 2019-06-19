package fulfiller

import (
	"testing"

	"github.com/Cepreu/gofrend/utils"
)

func TestGetScript(t *testing.T) {
	script := getScript("ivr_scripts/is_large_test.five9ivr")
	utils.PrettyPrint(script)
}

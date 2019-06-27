package utils

import (
	"testing"

	"github.com/Cepreu/gofrend/ivr"
)

func TestDisplayName(t *testing.T) {
	scriptHash := "1"
	moduleID := ivr.ModuleID("2")
	displayName := MakeDisplayName(scriptHash, moduleID)
	expected := "1-2"
	if displayName != expected {
		t.Fatalf("Unexpected display name. Expected: %s. Got: %s.", expected, displayName)
	}
}

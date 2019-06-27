package utils

import (
	"fmt"
	"strings"

	"github.com/Cepreu/gofrend/ivr"
)

// DisplayNameToScriptHash parses intent display name and returns script hash
func DisplayNameToScriptHash(intentName string) string {
	dashIndex := strings.LastIndex(intentName, "-")
	return intentName[:dashIndex]
}

// DisplayNameToModuleID parses intent display name and returns module ID
func DisplayNameToModuleID(intentName string) ivr.ModuleID {
	dashIndex := strings.LastIndex(intentName, "-")
	return ivr.ModuleID(intentName[dashIndex+1:])
}

// MakeDisplayName constructs intent display name from script hash and module ID
func MakeDisplayName(scriptHash string, moduleID ivr.ModuleID) string {
	return fmt.Sprintf("%s-%s", scriptHash, string(moduleID))
}

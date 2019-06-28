package utils

import (
	"fmt"
	"strings"

	"github.com/Cepreu/gofrend/ivr"
)

// DisplayNameToScriptHash parses intent display name and returns script hash
func DisplayNameToScriptHash(displayName string) string {
	dashIndex := strings.LastIndex(displayName, "-")
	return displayName[:dashIndex]
}

// DisplayNameToModuleID parses intent display name and returns module ID
func DisplayNameToModuleID(displayName string) ivr.ModuleID {
	dashIndex := strings.LastIndex(displayName, "-")
	return ivr.ModuleID(displayName[dashIndex+1:])
}

// MakeDisplayName constructs intent display name from script hash and module ID
func MakeDisplayName(scriptHash string, moduleID ivr.ModuleID) string {
	return fmt.Sprintf("%s-%s", scriptHash, string(moduleID))
}

// MakeContextName formats displayname correctly to be understood by dialoglflow as context
func MakeContextName(displayName string) string {
	return fmt.Sprintf("projects/f9-dialogflow-converter/agent/sessions/-/contexts/%s", displayName)
}

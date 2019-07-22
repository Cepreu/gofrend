package utils

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/Cepreu/gofrend/ivr"
)

// DisplayNameToScriptHash parses intent display name and returns script hash
func DisplayNameToScriptHash(displayName string) string {
	dashIndex := strings.LastIndex(displayName, "-")
	return displayName[:dashIndex]
}

// InputDisplayNameToModuleID parses intent display name for input module and returns module ID
func InputDisplayNameToModuleID(displayName string) ivr.ModuleID {
	dashIndex := strings.LastIndex(displayName, "-")
	return ivr.ModuleID(displayName[dashIndex+1:])
}

// MenuDisplayNameToModuleID parses intent display name for menu module and returns module ID
func MenuDisplayNameToModuleID(displayName string) ivr.ModuleID {
	dashIndex := strings.LastIndex(displayName, "-")
	dotIndex := strings.Index(displayName, ".")
	return ivr.ModuleID(displayName[dashIndex+1 : dotIndex])
}

// DisplayNameToModuleID parses display name and returns module ID
func DisplayNameToModuleID(displayName string) ivr.ModuleID {
	if strings.Contains(displayName, ".") {
		return MenuDisplayNameToModuleID(displayName)
	}
	return InputDisplayNameToModuleID(displayName)
}

// MenuDisplayNameToBranchName parses intent display name for menu module and returns branch name
func MenuDisplayNameToBranchName(displayName string) string {
	dotIndex := strings.Index(displayName, ".")
	hexName := displayName[dotIndex+1:]
	data, err := hex.DecodeString(hexName)
	if err != nil {
		panic(err)
	}
	return string(data)
}

// MakeInputDisplayName constructs intent display name for input module from script hash and module ID
func MakeInputDisplayName(scriptHash string, moduleID ivr.ModuleID) string {
	return fmt.Sprintf("%s-%s", scriptHash, string(moduleID))
}

// MakeMenuDisplayName constructs intent display name for menu module from script hash, module ID, and branch name
func MakeMenuDisplayName(scriptHash string, moduleID ivr.ModuleID, branchName string) string {
	return fmt.Sprintf("%s-%s.%s", scriptHash, string(moduleID), hex.EncodeToString([]byte(branchName)))
}

// MakeInputContextName formats input displayname correctly to be understood by dialoglflow as context
func MakeInputContextName(displayName string) string {
	return fmt.Sprintf("projects/f9-dialogflow-converter/agent/sessions/-/contexts/%s", displayName)
}

// MakeMenuContextName formats menu displayname correctly to be understood by dialoglflow as context
func MakeMenuContextName(displayName string) string {
	return fmt.Sprintf("projects/f9-dialogflow-converter/agent/sessions/-/contexts/%s", displayName[:strings.LastIndex(displayName, ".")])
}

// MenuDisplayNameToInputDisplayName removes the branch name from a menu display name
func MenuDisplayNameToInputDisplayName(displayName string) string {
	return displayName[:strings.LastIndex(displayName, ".")]
}

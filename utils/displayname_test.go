package utils

import (
	"testing"

	"github.com/Cepreu/gofrend/ivr"
)

func TestInputDisplayName(t *testing.T) {
	scriptHash := "1"
	moduleID := ivr.ModuleID("2")
	inputDisplayName := MakeInputDisplayName(scriptHash, moduleID)
	expected := "1-2"
	if inputDisplayName != expected {
		t.Fatalf("Unexpected display name. Expected: %s. Got: %s.", expected, inputDisplayName)
	}
	retrievedScriptHash := DisplayNameToScriptHash(inputDisplayName)
	if retrievedScriptHash != scriptHash {
		t.Fatalf("Unexpected script hash. Expected: %s. Got: %s.", scriptHash, retrievedScriptHash)
	}
	retrievedModuleID := InputDisplayNameToModuleID(inputDisplayName)
	if retrievedModuleID != moduleID {
		t.Fatalf("Unexpected moduleID. Expected: %s. Got: %s.", moduleID, retrievedModuleID)
	}
}

func TestMenuDisplayName(t *testing.T) {
	scriptHash := "1"
	moduleID := ivr.ModuleID("2")
	branchName := "3"
	menuDisplayName := MakeMenuDisplayName(scriptHash, moduleID, branchName)
	expected := "1-2.33"
	if menuDisplayName != expected {
		t.Fatalf("Unexpected display name. Expected: %s. Got: %s.", expected, menuDisplayName)
	}
	retrievedScriptHash := DisplayNameToScriptHash(menuDisplayName)
	if retrievedScriptHash != scriptHash {
		t.Fatalf("Unexpected script hash. Expected: %s. Got: %s.", scriptHash, retrievedScriptHash)
	}
	retrievedModuleID := MenuDisplayNameToModuleID(menuDisplayName)
	if retrievedModuleID != moduleID {
		t.Fatalf("Unexpected moduleID. Expected: %s. Got: %s.", moduleID, retrievedModuleID)
	}
	retrievedBranchName := MenuDisplayNameToBranchName(menuDisplayName)
	if retrievedBranchName != branchName {
		t.Fatalf("Uncexpected branch name. Expected: %s. Got: %s.", branchName, retrievedBranchName)
	}
}

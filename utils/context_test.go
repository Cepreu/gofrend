package utils

import (
	"testing"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func TestContext(t *testing.T) {
	projectID := "1"
	hash := "2"
	contextID := HashToContext(projectID, hash)
	expectedID := "projects/1/agent/sessions/-/contexts/2"
	if contextID != "projects/1/agent/sessions/-/contexts/2" {
		t.Fatalf("HashToContext produced incorrect context ID. Got: '%s'. Expected: '%s'.", contextID, expectedID)
	}
	retrievedHash, err := ContextToHash(contextID)
	if err != nil {
		t.Fatal(err)
	}
	if retrievedHash != hash {
		t.Fatalf("ContextToHash produced incorrect hash. Got: '%s'. Expected: '%s'.", retrievedHash, hash)
	}
}

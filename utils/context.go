package utils

import (
	"fmt"
	"strings"
)

// HashToContext formats projectID and hash strings correctly in the dialogflow context format
func HashToContext(projectID string, hash string) string {
	return fmt.Sprintf("projects/%s/agent/sessions/-/contexts/%s", projectID, hash)
}

// ContextToHash retrieves the hash from the context identifier
func ContextToHash(contextID string) (string, error) {
	lastIndex := strings.LastIndex(contextID, "/")
	if lastIndex < 0 || lastIndex == len(contextID)-1 {
		return "", fmt.Errorf("Could not parse context ID")
	}
	return contextID[lastIndex+1:], nil
}

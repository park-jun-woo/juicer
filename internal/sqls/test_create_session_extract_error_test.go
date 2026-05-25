//ff:func feature=ratchet type=session control=sequence
//ff:what TestCreateSession_ExtractError 테스트
package sqls

import (
	"testing"
)

func TestCreateSession_ExtractError(t *testing.T) {
	setupSessionDir(t)
	// Use a non-existent directory
	err := createSession("/nonexistent", "/nonexistent")
	if err == nil {
		t.Error("expected error for non-existent repo dir")
	}
}

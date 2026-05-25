//ff:func feature=hurl type=parse control=sequence
//ff:what TestRunStatus_WithSession 테스트
package hurls

import "testing"

func TestRunStatus_WithSession(t *testing.T) {
	cleanup := setupHurlTestSession(t)
	defer cleanup()
	if err := RunStatus(); err != nil {
		t.Fatal(err)
	}
}

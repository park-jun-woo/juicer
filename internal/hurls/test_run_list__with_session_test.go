//ff:func feature=hurl type=parse control=sequence
//ff:what TestRunList_WithSession 테스트
package hurls

import "testing"

func TestRunList_WithSession(t *testing.T) {
	cleanup := setupHurlTestSession(t)
	defer cleanup()
	if err := RunList(); err != nil {
		t.Fatal(err)
	}
}

//ff:func feature=hurl type=session control=sequence
//ff:what TestSessionPath 테스트
package hurls

import (
	"strings"
	"testing"
)

func TestSessionPath(t *testing.T) {
	p := sessionPath()
	if !strings.Contains(p, sessionFileName) {
		t.Fatalf("expected path to contain %s, got %s", sessionFileName, p)
	}
}

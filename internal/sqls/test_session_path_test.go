//ff:func feature=sql type=session control=sequence
//ff:what TestSessionPath 테스트
package sqls

import (
	"path/filepath"
	"testing"
)

func TestSessionPath(t *testing.T) {
	got := sessionPath()
	expected := filepath.Join(".codist", "sql-session.json")
	if got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

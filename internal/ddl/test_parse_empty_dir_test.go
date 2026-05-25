//ff:func feature=ddl type=parse control=sequence
//ff:what TestParse_EmptyDir 테스트
package ddl

import (
	"testing"
)

func TestParse_EmptyDir(t *testing.T) {
	dir := t.TempDir()
	tables, err := Parse(dir)
	if err != nil {
		t.Fatalf("Parse() error: %v", err)
	}
	if len(tables) != 0 {
		t.Errorf("expected 0 tables, got %d", len(tables))
	}
}

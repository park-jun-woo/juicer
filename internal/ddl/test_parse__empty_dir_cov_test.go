//ff:func feature=ddl type=parse control=sequence
//ff:what TestParse_EmptyDirCov 테스트
package ddl

import "testing"

func TestParse_EmptyDirCov(t *testing.T) {
	dir := t.TempDir()
	tables, err := Parse(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(tables) != 0 {
		t.Fatalf("expected 0 tables, got %d", len(tables))
	}
}

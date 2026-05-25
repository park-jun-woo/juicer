//ff:func feature=sql type=parse control=sequence
//ff:what TestExtractTables_Empty 테스트
package sqls

import "testing"

func TestExtractTables_Empty(t *testing.T) {
	tables := extractTables(nil)
	if len(tables) != 0 {
		t.Fatalf("expected 0, got %d", len(tables))
	}
}

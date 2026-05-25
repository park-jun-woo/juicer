//ff:func feature=sql type=parse control=sequence
//ff:what TestExtractTables_Basic 테스트
package sqls

import "testing"

func TestExtractTables_Basic(t *testing.T) {
	tables := extractTables([]string{"SELECT * FROM users"})
	if len(tables) == 0 {
		t.Fatal("expected at least one table")
	}
}

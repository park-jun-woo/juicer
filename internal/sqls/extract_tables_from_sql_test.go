//ff:func feature=sql type=parse control=iteration dimension=1
//ff:what TestExtractTablesFromSQL_Select 테스트
package sqls

import "testing"

func TestExtractTablesFromSQL_Select(t *testing.T) {
	tables := extractTablesFromSQL("SELECT * FROM users WHERE id = $1")
	found := false
	for _, tbl := range tables {
		if tbl == "users" {
			found = true
		}
	}
	if !found {
		t.Fatalf("expected users table, got %v", tables)
	}
}

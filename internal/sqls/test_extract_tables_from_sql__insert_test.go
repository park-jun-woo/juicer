//ff:func feature=sql type=parse control=iteration dimension=1
//ff:what TestExtractTablesFromSQL_Insert 테스트
package sqls

import "testing"

func TestExtractTablesFromSQL_Insert(t *testing.T) {
	tables := extractTablesFromSQL("INSERT INTO orders (id) VALUES ($1)")
	found := false
	for _, tbl := range tables {
		if tbl == "orders" {
			found = true
		}
	}
	if !found {
		t.Fatalf("expected orders table, got %v", tables)
	}
}

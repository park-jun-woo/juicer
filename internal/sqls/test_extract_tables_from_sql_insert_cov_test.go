//ff:func feature=sql type=test control=iteration dimension=1
//ff:what TestExtractTablesFromSQL_InsertCov 테스트
package sqls

import "testing"

func TestExtractTablesFromSQL_InsertCov(t *testing.T) {
	tables := extractTablesFromSQL("INSERT INTO orders (id) VALUES (1)")
	found := false
	for _, tbl := range tables {
		if tbl == "orders" {
			found = true
		}
	}
	if !found {
		t.Fatalf("expected orders, got %v", tables)
	}
}

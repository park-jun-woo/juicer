//ff:func feature=sql type=test control=iteration dimension=1
//ff:what TestExtractTablesFromSQL_ReservedWordCov 테스트
package sqls

import "testing"

func TestExtractTablesFromSQL_ReservedWordCov(t *testing.T) {
	// "select" after FROM should be treated as reserved word if it matches
	tables := extractTablesFromSQL("SELECT * FROM select WHERE 1=1")
	for _, tbl := range tables {
		if tbl == "select" {
			t.Fatal("should not include reserved word 'select'")
		}
	}
}

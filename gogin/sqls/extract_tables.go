//ff:func feature=sql type=parse control=iteration dimension=1
//ff:what SQL 조각들에서 정규식으로 테이블명 추출
package sqls

// extractTables extracts table names from SQL fragments using regex patterns.
//
func extractTables(fragments []string) []string {
	var tables []string
	for _, frag := range fragments {
		for _, t := range extractTablesFromSQL(frag) {
			tables = appendUnique(tables, t)
		}
	}
	return tables
}


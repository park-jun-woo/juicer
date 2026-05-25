//ff:func feature=sql type=parse control=iteration dimension=2
//ff:what 단일 SQL 문자열에서 테이블명 추출
package sqls

import (
	"regexp"
	"strings"
)

// extractTablesFromSQL extracts table names from a single SQL string.
//
func extractTablesFromSQL(sql string) []string {
	var tables []string
	for _, re := range []*regexp.Regexp{reFrom, reInsert, reUpdate, reDelete, reJoin} {
		matches := re.FindAllStringSubmatch(sql, -1)
		for _, m := range matches {
			tbl := strings.ToLower(m[1])
			if reservedWords[tbl] {
				continue
			}
			tables = appendUnique(tables, tbl)
		}
	}
	return tables
}


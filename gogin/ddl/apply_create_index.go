//ff:func feature=ddl type=parse control=sequence
//ff:what CREATE INDEX 를 대상 테이블에 등록
package ddl

import "strings"

// applyCreateIndex registers an index for the target table.
func applyCreateIndex(tables map[string]*Table, tableName, stmt string) {
	tableName = strings.ToLower(tableName)
	t := tables[tableName]
	if t == nil {
		// Table not yet created or already dropped; ignore
		return
	}
	t.Indexes = append(t.Indexes, stmt)
}

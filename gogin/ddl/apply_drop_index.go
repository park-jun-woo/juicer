//ff:func feature=ddl type=parse control=iteration dimension=2
//ff:what DROP INDEX 를 전체 테이블에서 해당 인덱스 제거
package ddl

import "strings"

// applyDropIndex removes an index by name from all tables.
func applyDropIndex(tables map[string]*Table, indexName string) {
	indexName = strings.ToLower(indexName)
	for _, t := range tables {
		filtered := t.Indexes[:0]
		for _, idx := range t.Indexes {
			if !strings.Contains(strings.ToLower(idx), indexName) {
				filtered = append(filtered, idx)
			}
		}
		t.Indexes = filtered
	}
}

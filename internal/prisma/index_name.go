//ff:func feature=prisma type=convert control=sequence topic=prisma
//ff:what 테이블명+컬럼들로 인덱스 이름(table_col1_col2_idx) 생성
package prisma

import "strings"

// indexName builds a deterministic index name like table_col1_col2_idx.
func indexName(table string, cols []string) string {
	parts := make([]string, 0, len(cols)+2)
	parts = append(parts, table)
	parts = append(parts, cols...)
	parts = append(parts, "idx")
	return strings.Join(parts, "_")
}

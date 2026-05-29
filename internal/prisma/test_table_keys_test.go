//ff:func feature=prisma type=test control=iteration dimension=1
//ff:what 테스트 헬퍼: 테이블 맵의 키 목록 반환
package prisma

import "github.com/park-jun-woo/codistill/internal/ddl"

// tableKeys returns the sorted-agnostic key list of a parsed table map.
func tableKeys(tables map[string]*ddl.Table) []string {
	keys := make([]string, 0, len(tables))
	for k := range tables {
		keys = append(keys, k)
	}
	return keys
}

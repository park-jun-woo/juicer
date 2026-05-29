//ff:func feature=prisma type=test control=iteration dimension=1
//ff:what 테스트 헬퍼: 컬럼 목록에서 이름으로 컬럼 검색
package prisma

import "github.com/park-jun-woo/codistill/internal/ddl"

// findColumn returns the column with the given name, or nil if absent.
func findColumn(cols []ddl.Column, name string) *ddl.Column {
	for i := range cols {
		if cols[i].Name == name {
			return &cols[i]
		}
	}
	return nil
}

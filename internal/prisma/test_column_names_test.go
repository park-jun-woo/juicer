//ff:func feature=prisma type=test control=iteration dimension=1
//ff:what 테스트 헬퍼: 컬럼 목록의 이름 목록 반환
package prisma

import "github.com/park-jun-woo/codistill/internal/ddl"

// columnNames returns the names of all columns for diagnostic output.
func columnNames(cols []ddl.Column) []string {
	names := make([]string, 0, len(cols))
	for _, c := range cols {
		names = append(names, c.Name)
	}
	return names
}

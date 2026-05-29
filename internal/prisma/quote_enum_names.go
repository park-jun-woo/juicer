//ff:func feature=prisma type=convert control=iteration dimension=1 topic=prisma
//ff:what enum 타입 목록의 이름만 인용한 새 목록 반환 (정의 출력이 컬럼 타입 위치와 동일 문자열)
package prisma

import "github.com/park-jun-woo/codistill/internal/ddl"

// quoteEnumNames returns a copy of enums with each type name quoted, so the
// CREATE TYPE definition emits the same quoted string ("Role") used at the
// column type position. The originals stay raw for schema enum-set lookups.
func quoteEnumNames(enums []ddl.EnumType) []ddl.EnumType {
	out := make([]ddl.EnumType, 0, len(enums))
	for _, e := range enums {
		out = append(out, ddl.EnumType{Name: quoteIdent(e.Name), Values: e.Values})
	}
	return out
}

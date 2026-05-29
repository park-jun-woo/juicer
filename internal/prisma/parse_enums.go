//ff:func feature=prisma type=parse control=iteration dimension=1 topic=prisma
//ff:what Prisma 소스 문자열을 주석제거 후 enum 정의 목록(ddl.EnumType)으로 파싱
package prisma

import "github.com/park-jun-woo/codistill/internal/ddl"

// parseEnums strips comments and parses all enum blocks in one source string.
func parseEnums(src string) []ddl.EnumType {
	blocks := findEnumBlocks(stripComments(src))
	enums := make([]ddl.EnumType, 0, len(blocks))
	for name, body := range blocks {
		enums = append(enums, ddl.EnumType{Name: name, Values: parseEnumValues(body)})
	}
	return enums
}

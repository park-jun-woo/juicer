//ff:type feature=prisma type=model topic=prisma
//ff:what 파싱된 model <Name> { ... } 블록의 중간 표현 구조체
package prisma

// model is one parsed `model <Name> { ... }` block before DDL conversion.
type model struct {
	name       string // Prisma model name
	tableName  string // @@map override or model name
	fields     []field
	blockAttrs []string // raw @@... block attribute lines
}

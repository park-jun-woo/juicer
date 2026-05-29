//ff:type feature=prisma type=model topic=prisma
//ff:what 모델 블록 내 한 라인을 파싱한 필드 중간 표현 구조체
package prisma

// field is one parsed line inside a model block.
type field struct {
	name     string   // Prisma field name
	baseType string   // type without ? or []
	nullable bool     // trailing ?
	array    bool     // trailing []
	attrs    []string // raw @... field attributes
}

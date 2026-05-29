//ff:func feature=prisma type=convert control=sequence topic=prisma
//ff:what 필드 base 타입이 선언된 model 집합에 속하면 관계 필드로 판정
package prisma

// isRelationField reports whether f references another declared model
// (a navigation field, not a stored column).
func isRelationField(f field, models map[string]bool) bool {
	return models[f.baseType]
}

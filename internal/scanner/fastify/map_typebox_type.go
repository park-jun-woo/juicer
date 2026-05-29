//ff:func feature=scan type=extract control=selection topic=fastify
//ff:what TypeBox Type.X 메서드명을 OpenAPI 스칼라 타입 문자열로 매핑한다 (복합/미지원은 빈 문자열)
package fastify

func mapTypeBoxType(name string) string {
	switch name {
	case "String":
		return "string"
	case "Number":
		return "number"
	case "Integer":
		return "integer"
	case "Boolean":
		return "boolean"
	default:
		return ""
	}
}

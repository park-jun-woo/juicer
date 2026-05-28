//ff:func feature=scan type=extract control=selection topic=fastify
//ff:what JSON Schema 타입 문자열을 Go 타입명으로 매핑한다
package fastify

func mapJSONSchemaType(t string) string {
	switch t {
	case "string":
		return "string"
	case "integer":
		return "integer"
	case "number":
		return "number"
	case "boolean":
		return "boolean"
	case "array":
		return "array"
	case "object":
		return "object"
	default:
		return t
	}
}

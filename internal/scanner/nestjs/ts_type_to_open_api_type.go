//ff:func feature=scan type=convert control=sequence topic=nestjs
//ff:what TypeScript 타입에서 OpenAPI 타입 문자열만 반환한다
package nestjs

// tsTypeToOpenAPIType returns the simple OpenAPI type string for a TypeScript type.
func tsTypeToOpenAPIType(ts string) string {
	t := tsTypeToOpenAPI(ts)
	if t.Type == "" {
		return "string"
	}
	if t.Format != "" {
		return t.Type + ":" + t.Format
	}
	return t.Type
}

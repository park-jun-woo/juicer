//ff:func feature=scan type=convert control=sequence topic=fastapi
//ff:what Python 타입을 OpenAPI 타입 문자열로 변환한다
package fastapi

// mapTypeToOpenAPI converts a Python type to an OpenAPI type string.
func mapTypeToOpenAPI(typeName string) string {
	oa := pyTypeToOpenAPI(typeName)
	if oa.Type == "" {
		return "string"
	}
	if oa.Format != "" {
		return oa.Type + ":" + oa.Format
	}
	return oa.Type
}

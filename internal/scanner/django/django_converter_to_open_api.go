//ff:func feature=scan type=convert control=selection topic=django
//ff:what Django URL 컨버터 타입을 OpenAPI 타입으로 변환한다
package django

// djangoConverterToOpenAPI converts a Django URL converter type to OpenAPI type.
func djangoConverterToOpenAPI(converter string) openAPIType {
	switch converter {
	case "int":
		return openAPIType{Type: "integer"}
	case "uuid":
		return openAPIType{Type: "string", Format: "uuid"}
	case "slug":
		return openAPIType{Type: "string"}
	case "path":
		return openAPIType{Type: "string"}
	case "str", "":
		return openAPIType{Type: "string"}
	default:
		return openAPIType{Type: "string"}
	}
}

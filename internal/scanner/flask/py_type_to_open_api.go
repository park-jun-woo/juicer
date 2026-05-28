//ff:func feature=scan type=convert control=selection topic=flask
//ff:what Flask URL 컨버터 타입을 OpenAPI 타입으로 변환한다
package flask

// flaskConverterToOpenAPI converts a Flask URL converter type to OpenAPI type.
// Flask converters: string (default), int, float, path, uuid.
func flaskConverterToOpenAPI(converter string) openAPIType {
	switch converter {
	case "int":
		return openAPIType{Type: "integer", Format: "int64"}
	case "float":
		return openAPIType{Type: "number", Format: "double"}
	case "uuid":
		return openAPIType{Type: "string", Format: "uuid"}
	case "path":
		return openAPIType{Type: "string"}
	case "string", "":
		return openAPIType{Type: "string"}
	default:
		return openAPIType{Type: "string"}
	}
}

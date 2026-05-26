//ff:func feature=scan type=convert control=selection topic=fastapi
//ff:what Python 타입을 OpenAPI 타입으로 변환한다
package fastapi

import "strings"

// pyTypeToOpenAPI converts a Python type string to an openAPIType.
func pyTypeToOpenAPI(py string) openAPIType {
	py = strings.TrimSpace(py)

	if result, ok := tryNullable(py); ok {
		return result
	}
	if result, ok := tryList(py); ok {
		return result
	}

	switch py {
	case "str":
		return openAPIType{Type: "string"}
	case "int":
		return openAPIType{Type: "integer"}
	case "float":
		return openAPIType{Type: "number"}
	case "bool":
		return openAPIType{Type: "boolean"}
	case "dict":
		return openAPIType{Type: "object"}
	case "datetime":
		return openAPIType{Type: "string", Format: "date-time"}
	case "date":
		return openAPIType{Type: "string", Format: "date"}
	case "EmailStr":
		return openAPIType{Type: "string", Format: "email"}
	case "Any":
		return openAPIType{Type: "object"}
	case "":
		return openAPIType{}
	default:
		return openAPIType{Type: "object"}
	}
}

//ff:func feature=scan type=convert control=selection topic=dotnet
//ff:what C# 타입을 OpenAPI 타입으로 변환한다
package dotnet

import "strings"

func csharpTypeToOpenAPI(ctype string) openAPIType {
	ctype = strings.TrimSpace(ctype)
	ctype = strings.TrimSuffix(ctype, "?")

	switch ctype {
	case "string":
		return openAPIType{Type: "string"}
	case "int", "Int32":
		return openAPIType{Type: "integer", Format: "int32"}
	case "long", "Int64":
		return openAPIType{Type: "integer", Format: "int64"}
	case "short", "Int16":
		return openAPIType{Type: "integer", Format: "int32"}
	case "float", "Single":
		return openAPIType{Type: "number", Format: "float"}
	case "double", "Double":
		return openAPIType{Type: "number", Format: "double"}
	case "decimal", "Decimal":
		return openAPIType{Type: "number"}
	case "bool", "Boolean":
		return openAPIType{Type: "boolean"}
	case "DateTime", "DateTimeOffset":
		return openAPIType{Type: "string", Format: "date-time"}
	case "DateOnly":
		return openAPIType{Type: "string", Format: "date"}
	case "Guid":
		return openAPIType{Type: "string", Format: "uuid"}
	case "IFormFile":
		return openAPIType{Type: "string", Format: "binary"}
	case "byte[]":
		return openAPIType{Type: "string", Format: "byte"}
	case "void":
		return openAPIType{}
	}

	if strings.HasPrefix(ctype, "List<") || strings.HasPrefix(ctype, "IEnumerable<") ||
		strings.HasPrefix(ctype, "IList<") || strings.HasPrefix(ctype, "ICollection<") {
		inner := extractGenericInner(ctype)
		return openAPIType{Type: "array", Items: inner}
	}

	if strings.HasPrefix(ctype, "Dictionary<") || strings.HasPrefix(ctype, "IDictionary<") {
		return openAPIType{Type: "object"}
	}

	if strings.HasSuffix(ctype, "[]") {
		inner := ctype[:len(ctype)-2]
		return openAPIType{Type: "array", Items: inner}
	}

	return openAPIType{Type: "object"}
}

//ff:func feature=scan type=convert control=selection topic=spring
//ff:what Java 타입을 OpenAPI 타입으로 변환한다
package spring

import "strings"

func javaTypeToOpenAPI(jtype string) openAPIType {
	jtype = strings.TrimSpace(jtype)

	switch jtype {
	case "String":
		return openAPIType{Type: "string"}
	case "int", "Integer":
		return openAPIType{Type: "integer", Format: "int32"}
	case "long", "Long":
		return openAPIType{Type: "integer", Format: "int64"}
	case "float", "Float":
		return openAPIType{Type: "number", Format: "float"}
	case "double", "Double":
		return openAPIType{Type: "number", Format: "double"}
	case "boolean", "Boolean":
		return openAPIType{Type: "boolean"}
	case "BigDecimal":
		return openAPIType{Type: "number"}
	case "LocalDateTime", "ZonedDateTime", "Instant", "OffsetDateTime":
		return openAPIType{Type: "string", Format: "date-time"}
	case "LocalDate":
		return openAPIType{Type: "string", Format: "date"}
	case "UUID":
		return openAPIType{Type: "string", Format: "uuid"}
	case "MultipartFile":
		return openAPIType{Type: "string", Format: "binary"}
	case "byte[]":
		return openAPIType{Type: "string", Format: "byte"}
	case "void", "Void", "":
		return openAPIType{}
	}

	if isCollectionType(jtype) {
		inner := extractGenericInner(jtype)
		return openAPIType{Type: "array", Items: inner}
	}

	if strings.HasPrefix(jtype, "Map<") || strings.HasPrefix(jtype, "HashMap<") {
		return openAPIType{Type: "object"}
	}

	if strings.HasSuffix(jtype, "[]") {
		inner := jtype[:len(jtype)-2]
		return openAPIType{Type: "array", Items: inner}
	}

	return openAPIType{Type: "object"}
}

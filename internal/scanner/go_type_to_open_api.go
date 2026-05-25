//ff:func feature=scan type=extract control=selection
//ff:what Go 타입 문자열을 OpenAPI type으로 변환한다
package scanner

import (
	"strings"
)

func goTypeToOpenAPI(goType string) string {
	// 포인터 unwrap
	if strings.HasPrefix(goType, "*") {
		goType = goType[1:]
	}

	switch goType {
	case "string":
		return "string"
	case "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64":
		return "integer"
	case "float32", "float64":
		return "number"
	case "bool", "boolean":
		return "boolean"
	case "integer":
		return "integer"
	case "number":
		return "number"
	case "object":
		return "object"
	case "array":
		return "array"
	case "any", "interface{}":
		return "object"
	}

	// 슬라이스
	if strings.HasPrefix(goType, "[]") {
		return "array"
	}

	// time.Time 등 알려진 타입
	if goType == "time.Time" {
		return "string"
	}

	return "object"
}


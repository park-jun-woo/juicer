//ff:func feature=scan type=extract control=sequence
//ff:what 단일 Field를 OpenAPI property로 변환한다
package scanner

import (
	"strings"
)

func fieldToProperty(f Field) map[string]any {
	goType := f.Type
	prop := map[string]any{}

	// 중첩 struct
	if len(f.Fields) > 0 {
		if strings.HasPrefix(goType, "[]") {
			prop["type"] = "array"
			prop["items"] = fieldsToSchema(f.Fields)
			return prop
		}
		return fieldsToSchema(f.Fields)
	}

	// 배열 타입
	if strings.HasPrefix(goType, "[]") {
		elemType := goType[2:]
		prop["type"] = "array"
		prop["items"] = map[string]any{"type": goTypeToOpenAPI(elemType)}
		return prop
	}

	// 포인터 unwrap
	if strings.HasPrefix(goType, "*") {
		goType = goType[1:]
	}

	oaType := goTypeToOpenAPI(goType)
	prop["type"] = oaType

	// format 힌트
	if format := goTypeFormat(goType, f); format != "" {
		prop["format"] = format
	}

	return prop
}


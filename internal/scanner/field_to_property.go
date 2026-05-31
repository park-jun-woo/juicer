//ff:func feature=scan type=extract control=sequence
//ff:what 단일 Field를 OpenAPI property로 변환한다
package scanner

import (
	"strings"
)

func fieldToProperty(f Field) map[string]any {
	goType := f.Type
	prop := map[string]any{}

	if f.Nullable {
		prop["nullable"] = true
	}

	// named 타입 참조 ($ref). 배열이면 type:array + items:$ref.
	if f.Ref != "" {
		ref := map[string]any{"$ref": "#/components/schemas/" + f.Ref}
		if goType == "array" || strings.HasPrefix(goType, "[]") {
			prop["type"] = "array"
			prop["items"] = ref
			return prop
		}
		return ref
	}

	// 중첩 struct
	if len(f.Fields) > 0 {
		if strings.HasPrefix(goType, "[]") {
			prop["type"] = "array"
			prop["items"] = fieldsToSchema(f.Fields)
			return prop
		}
		schema := fieldsToSchema(f.Fields)
		if f.Nullable {
			schema["nullable"] = true
		}
		return schema
	}

	// 배열 타입
	if strings.HasPrefix(goType, "[]") {
		elemType := goType[2:]
		prop["type"] = "array"
		oaElem := goTypeToOpenAPI(elemType)
		if oaElem == "object" {
			prop["items"] = map[string]any{"$ref": "#/components/schemas/" + elemType}
		} else {
			prop["items"] = map[string]any{"type": oaElem}
		}
		return prop
	}

	// 포인터 unwrap
	if strings.HasPrefix(goType, "*") {
		goType = goType[1:]
	}

	// type:format 규칙 (e.g. "string:date-time")
	if i := strings.IndexByte(goType, ':'); i >= 0 {
		prop["type"] = goType[:i]
		prop["format"] = goType[i+1:]
		return prop
	}

	oaType := goTypeToOpenAPI(goType)
	prop["type"] = oaType

	// format 힌트
	if format := goTypeFormat(goType, f); format != "" {
		prop["format"] = format
	}

	// enum 값
	if len(f.Enum) > 0 {
		prop["enum"] = f.Enum
	}

	// 제약조건
	if f.Minimum != nil {
		prop["minimum"] = *f.Minimum
	}
	if f.Maximum != nil {
		prop["maximum"] = *f.Maximum
	}
	if f.MinLength != nil {
		prop["minLength"] = *f.MinLength
	}
	if f.MaxLength != nil {
		prop["maxLength"] = *f.MaxLength
	}

	return prop
}


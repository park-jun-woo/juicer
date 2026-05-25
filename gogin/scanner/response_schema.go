//ff:func feature=scan type=extract control=sequence
//ff:what 응답의 named/anonymous 여부에 따라 $ref 또는 inline schema를 반환한다
package scanner

import (
	"strings"
)

func responseSchema(resp Response, schemas map[string]any) map[string]any {
	// gin.H는 named type이지만 components에 등록하지 않음
	if resp.TypeName != "" && resp.TypeName != "gin.H" && len(resp.Fields) > 0 {
		typeName := resp.TypeName
		isSlice := strings.HasPrefix(typeName, "[]")
		if isSlice {
			typeName = typeName[2:]
		}
		schemaName := lcFirst(typeName)
		if _, exists := schemas[schemaName]; !exists {
			schemas[schemaName] = fieldsToSchema(resp.Fields)
		}
		ref := map[string]any{"$ref": "#/components/schemas/" + schemaName}
		if isSlice {
			return map[string]any{"type": "array", "items": ref}
		}
		return ref
	}

	if len(resp.Fields) > 0 {
		return fieldsToSchema(resp.Fields)
	}

	if resp.TypeName != "" {
		return map[string]any{"type": "object"}
	}

	return map[string]any{"type": "object"}
}


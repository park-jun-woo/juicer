//ff:func feature=scan type=extract control=sequence
//ff:what Body의 named/anonymous 여부에 따라 $ref 또는 inline schema를 반환한다
package scanner

import (
	"strings"
)

func bodySchema(body *Body, schemas map[string]any) map[string]any {
	if body.TypeName != "" && len(body.Fields) > 0 {
		typeName := body.TypeName
		isSlice := strings.HasPrefix(typeName, "[]")
		if isSlice {
			typeName = typeName[2:]
		}
		schemaName := lcFirst(typeName)
		if _, exists := schemas[schemaName]; !exists {
			schemas[schemaName] = fieldsToSchema(body.Fields)
		}
		ref := map[string]any{"$ref": "#/components/schemas/" + schemaName}
		if isSlice {
			return map[string]any{"type": "array", "items": ref}
		}
		return ref
	}

	if len(body.Fields) > 0 {
		return fieldsToSchema(body.Fields)
	}

	return map[string]any{"type": "object"}
}


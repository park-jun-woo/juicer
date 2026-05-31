//ff:func feature=scan type=extract control=sequence
//ff:what 응답의 named/anonymous 여부에 따라 $ref 또는 inline schema를 반환한다
package scanner

func responseSchema(resp Response, schemas map[string]any) map[string]any {
	// gin.H는 named type이지만 components에 등록하지 않음
	if resp.TypeName != "" && resp.TypeName != "gin.H" {
		primSchema, baseName, isSlice := resolvePrimitiveSchema(resp.TypeName)
		if primSchema != nil {
			return primSchema
		}

		schemaName := baseName // preserve original class casing (e.g. UpperCamel)
		ensureSchema(schemaName, resp.Fields, schemas)
		ref := map[string]any{"$ref": "#/components/schemas/" + schemaName}
		if isSlice {
			return map[string]any{"type": "array", "items": ref}
		}
		return ref
	}

	if len(resp.Fields) > 0 {
		return fieldsToSchema(resp.Fields)
	}

	return map[string]any{"type": "object"}
}


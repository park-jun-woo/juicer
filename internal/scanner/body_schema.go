//ff:func feature=scan type=extract control=sequence
//ff:what Body의 named/anonymous 여부에 따라 $ref 또는 inline schema를 반환한다
package scanner

func bodySchema(body *Body, schemas map[string]any) map[string]any {
	if body.TypeName != "" {
		primSchema, baseName, isSlice := resolvePrimitiveSchema(body.TypeName)
		if primSchema != nil {
			return primSchema
		}

		schemaName := baseName // preserve original class casing (e.g. UpperCamel)
		ensureSchema(schemaName, body.Fields, schemas)
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


//ff:func feature=scan type=extract control=sequence
//ff:what schemas에 schemaName이 없으면 fields 기반 또는 빈 object schema를 등록한다
package scanner

func ensureSchema(schemaName string, fields []Field, schemas map[string]any) {
	if _, exists := schemas[schemaName]; exists {
		return
	}
	if len(fields) > 0 {
		schemas[schemaName] = fieldsToSchema(fields)
		return
	}
	schemas[schemaName] = map[string]any{"type": "object"}
}

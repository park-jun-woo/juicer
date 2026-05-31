//ff:func feature=scan type=convert control=sequence
//ff:what Field 목록을 OpenAPI schema object로 변환한다(외부 스캐너용 공개 래퍼)
package scanner

// SchemaFromFields converts a list of Fields into an OpenAPI schema object.
// It is the exported wrapper around fieldsToSchema, used by framework scanners
// (e.g. NestJS) to register recursively-extracted named DTO schemas into
// ScanResult.Schemas.
func SchemaFromFields(fields []Field) map[string]any {
	return fieldsToSchema(fields)
}

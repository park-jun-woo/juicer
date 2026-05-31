//ff:func feature=scan type=convert control=sequence
//ff:what 문자열 enum 값 목록을 OpenAPI string-enum schema object로 변환한다
package scanner

// EnumSchema builds an OpenAPI string-enum schema object from the given values.
// Used by framework scanners (e.g. NestJS) to register enum component schemas.
func EnumSchema(values []string) map[string]any {
	return map[string]any{
		"type": "string",
		"enum": values,
	}
}

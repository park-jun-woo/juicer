//ff:func feature=scan type=convert control=sequence topic=fastapi
//ff:what 단일 pydanticField를 scanner.Field로 변환한다
package fastapi

import "github.com/park-jun-woo/juicer/internal/scanner"

// convertOnePydanticField converts a single pydanticField to scanner.Field.
func convertOnePydanticField(f pydanticField) scanner.Field {
	oa := pyTypeToOpenAPI(f.typeName)
	oaType := oa.Type
	if oaType == "" {
		oaType = "string"
	}
	if oa.Format != "" {
		oaType = oaType + ":" + oa.Format
	}
	return scanner.Field{
		Name:      f.name,
		Type:      oaType,
		JSON:      f.name,
		Nullable:  f.nullable || oa.Nullable,
		Minimum:   f.ge,
		Maximum:   f.le,
		MinLength: f.minLength,
		MaxLength: f.maxLength,
	}
}

//ff:func feature=scan type=parse control=selection topic=joi
//ff:what Joi 메서드를 Field에 반영한다 (type, format, required 등)
package joi

import "github.com/park-jun-woo/codistill/internal/scanner"

// ApplyMethod — 단일 ChainMethod를 Field에 반영.
// Joi는 기본 optional이며 .required()가 명시적이다. OpenAPI required 산출은
// scanner.isRequired가 Field.Validate의 "required" 문자열로 판정하므로
// .required()는 반드시 Validate에 "required"를 추가한다.
func ApplyMethod(f *scanner.Field, m ChainMethod) {
	switch m.Name {
	case "string":
		f.Type = "string"
	case "number":
		f.Type = "number"
	case "boolean":
		f.Type = "boolean"
	case "date":
		f.Type = "string"
		f.Validate = appendValidate(f.Validate, "date-time")
	case "integer":
		f.Type = "integer"
	case "array":
		f.Type = "array"
	case "object":
		f.Type = "object"
	case "required":
		f.Validate = appendValidate(f.Validate, "required")
	case "email":
		f.Validate = appendValidate(f.Validate, "email")
	case "uri":
		f.Validate = appendValidate(f.Validate, "uri")
	case "uuid", "guid":
		f.Validate = appendValidate(f.Validate, "uuid")
	case "valid":
		f.Type = "string"
		if len(m.Args) > 0 {
			f.Enum = m.Args
		}
	}
}

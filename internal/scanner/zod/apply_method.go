//ff:func feature=scan type=parse control=selection topic=zod
//ff:what Zod 메서드를 Field에 반영한다 (type, format, constraint 등)
package zod

import "github.com/park-jun-woo/codistill/internal/scanner"

// ApplyMethod — 단일 ChainMethod를 Field에 반영
func ApplyMethod(f *scanner.Field, m ChainMethod) {
	switch m.Name {
	case "string":
		f.Type = "string"
	case "number":
		f.Type = "number"
	case "boolean":
		f.Type = "boolean"
	case "int":
		f.Type = "integer"
	case "email":
		f.Validate = appendValidate(f.Validate, "email")
	case "url":
		f.Validate = appendValidate(f.Validate, "uri")
	case "uuid":
		f.Validate = appendValidate(f.Validate, "uuid")
	case "optional":
		f.Nullable = true
	case "nullable":
		f.Nullable = true
	case "min":
		ApplyMin(f, m)
	case "max":
		ApplyMax(f, m)
	case "enum":
		f.Type = "string"
		if len(m.Args) > 0 {
			f.Enum = m.Args
		}
	case "array":
		f.Type = "array"
	case "object":
		f.Type = "object"
	}
}

//ff:func feature=scan type=convert control=sequence topic=nestjs
//ff:what 단일 dtoField를 scanner.Field로 변환한다
package nestjs

import "github.com/park-jun-woo/codistill/internal/scanner"

// convertOneDtoField converts a single dtoField to scanner.Field.
func convertOneDtoField(f dtoField) scanner.Field {
	sf := tsTypeToField(f.name, f.tsType, f.optional)
	if f.validate != "" {
		sf.Validate = f.validate // 팩토리 경유: 원본 보존
	} else {
		sf.Validate = buildValidateTag(f.optional, f.validators)
	}
	if hasIsEnum(f.validators) {
		sf.Type = "string"
	}
	sf.Nullable = f.nullable
	sf.Enum = f.enum
	sf.MinLength = f.minLength
	sf.MaxLength = f.maxLength
	return sf
}

//ff:func feature=scan type=convert control=sequence topic=nestjs
//ff:what scanner.Field → dtoField 공통 변환 헬퍼
package nestjs

import (
	"strings"

	"github.com/park-jun-woo/juicer/internal/scanner"
)

// fieldToDTOField converts a scanner.Field to a dtoField, preserving all metadata.
func fieldToDTOField(f scanner.Field) dtoField {
	df := dtoField{
		name:      f.Name,
		tsType:    f.Type,
		validate:  f.Validate,
		nullable:  f.Nullable,
		enum:      f.Enum,
		minLength: f.MinLength,
		maxLength: f.MaxLength,
	}
	if df.tsType == "" {
		df.tsType = "string"
	}
	// Validate에 "optional"이 포함되어 있거나 "required"가 없으면 optional=true
	if strings.Contains(f.Validate, "optional") || !strings.Contains(f.Validate, "required") {
		df.optional = true
	}
	return df
}

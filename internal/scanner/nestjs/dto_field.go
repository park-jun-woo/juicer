//ff:type feature=scan type=model topic=nestjs
//ff:what DTO 클래스 필드 추출 결과 구조체
package nestjs

// dtoField represents a single field extracted from a DTO class.
type dtoField struct {
	name         string
	tsType       string
	optional     bool
	validators   []string
	enum         []string
	enumTypeName string // @IsEnum(X)의 X를 저장
	minLength    *int
	maxLength    *int
	validate     string // 원본 scanner.Field.Validate 보존
	nullable     bool   // 원본 scanner.Field.Nullable 보존
}

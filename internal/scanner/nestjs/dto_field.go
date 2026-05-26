//ff:type feature=scan type=model topic=nestjs
//ff:what DTO 클래스 필드 추출 결과 구조체
package nestjs

// dtoField represents a single field extracted from a DTO class.
type dtoField struct {
	name       string
	tsType     string
	optional   bool
	validators []string
}

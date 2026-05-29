//ff:type feature=scan type=model topic=express
//ff:what 크로스파일 Joi 검증 참조(`validate(authValidation.register)`) 구조체
package express

// joiValidatorRef — validate(<importName>.<member>) 형태의 크로스파일 참조.
// importName은 검증 스키마 파일의 import 바인딩명, member는 그 파일 내 const 이름.
type joiValidatorRef struct {
	ImportName string
	Member     string
}

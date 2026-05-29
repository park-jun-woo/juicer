//ff:type feature=scan type=model topic=joi
//ff:what Joi 메서드 체인의 단일 메서드 정보 구조체
package joi

// ChainMethod — Joi 메서드 체인의 단일 메서드 정보
type ChainMethod struct {
	Name string
	Args []string
}

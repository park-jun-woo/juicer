//ff:type feature=scan type=model topic=fastify
//ff:what async wrapper register의 본문 바이트 범위와 prefix를 담는 스코프 구조체
package fastify

type wrapperScope struct {
	Start  uint32
	End    uint32
	Prefix string
}

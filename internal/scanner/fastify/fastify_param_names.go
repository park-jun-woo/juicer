//ff:type feature=scan type=model topic=fastify
//ff:what Fastify 인스턴스로 인식되는 함수 파라미터 이름 집합
package fastify

var fastifyParamNames = map[string]bool{
	"fastify":  true,
	"app":      true,
	"server":   true,
	"instance": true,
}

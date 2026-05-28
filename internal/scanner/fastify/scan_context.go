//ff:type feature=scan type=model topic=fastify
//ff:what 스캔 컨텍스트 구조체 (Pass 1 결과)
package fastify

type scanContext struct {
	parsed    map[string]*fileInfo
	instances map[string]map[string]bool
	prefixMap map[string]string
	absRoot   string
}

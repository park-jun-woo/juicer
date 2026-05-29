//ff:type feature=scan type=model topic=fastify
//ff:what 추출된 라우트 정보 구조체
package fastify

import sitter "github.com/smacker/go-tree-sitter"

type routeInfo struct {
	Method    string
	Path      string
	Handler   string
	Line      int
	StartByte uint32
	Schema    *sitter.Node
}

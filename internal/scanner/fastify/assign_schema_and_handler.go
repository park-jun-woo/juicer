//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what 라우트 인자에서 schema 객체와 handler를 routeInfo에 할당한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func assignSchemaAndHandler(ri *routeInfo, argNodes []*sitter.Node, src []byte) {
	for i := 1; i < len(argNodes); i++ {
		assignOneArg(ri, argNodes[i], src)
	}
	if ri.Handler == "" {
		ri.Handler = "(anonymous)"
	}
}

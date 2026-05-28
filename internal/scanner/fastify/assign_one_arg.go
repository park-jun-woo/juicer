//ff:func feature=scan type=extract control=selection topic=fastify
//ff:what 단일 인자 노드를 routeInfo의 schema 또는 handler에 할당한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func assignOneArg(ri *routeInfo, arg *sitter.Node, src []byte) {
	switch arg.Type() {
	case "object":
		ri.Schema = arg
	case "identifier":
		ri.Handler = nodeText(arg, src)
	case "arrow_function", "function":
		ri.Handler = "(anonymous)"
	}
}

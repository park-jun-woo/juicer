//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what TypeBox call_expression의 arguments에서 첫 의미 있는 인자 노드(object 또는 중첩 call)를 반환한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func typeBoxFirstArg(call *sitter.Node) *sitter.Node {
	args := findChildByType(call, "arguments")
	if args == nil {
		return nil
	}
	for i := 0; i < int(args.ChildCount()); i++ {
		child := args.Child(i)
		if child.Type() == "object" || child.Type() == "call_expression" {
			return child
		}
	}
	return nil
}

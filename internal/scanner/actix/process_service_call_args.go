//ff:func feature=scan type=extract control=iteration dimension=1 topic=actix
//ff:what arguments 노드의 각 call_expression 인자를 서비스 인자로 처리한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func processServiceCallArgs(args *sitter.Node, src []byte, prefix string, routes *[]builderRoute) {
	for i := 0; i < int(args.ChildCount()); i++ {
		arg := args.Child(i)
		if arg.Type() == "call_expression" {
			processServiceArg(arg, src, prefix, routes)
		}
	}
}

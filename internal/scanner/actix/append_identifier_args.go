//ff:func feature=scan type=extract control=iteration dimension=1 topic=actix
//ff:what arguments 노드의 식별자 인자 이름들을 목록에 추가한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func appendIdentifierArgs(args *sitter.Node, src []byte, handlers []string) []string {
	for i := 0; i < int(args.ChildCount()); i++ {
		arg := args.Child(i)
		if arg.Type() == "identifier" {
			handlers = append(handlers, nodeText(arg, src))
		}
	}
	return handlers
}

//ff:func feature=scan type=extract control=iteration dimension=1 topic=actix
//ff:what arguments 노드에서 첫 표현식 인자(구두점 제외)를 반환한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func firstArgExpr(args *sitter.Node) *sitter.Node {
	for i := 0; i < int(args.ChildCount()); i++ {
		child := args.Child(i)
		if child.IsNamed() {
			return child
		}
	}
	return nil
}

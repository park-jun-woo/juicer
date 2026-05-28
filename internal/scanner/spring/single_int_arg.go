//ff:func feature=scan type=extract control=iteration dimension=1 topic=spring
//ff:what 어노테이션의 단일 정수 인자를 반환한다
package spring

import sitter "github.com/smacker/go-tree-sitter"

func singleIntArg(ann *sitter.Node, src []byte) (int, bool) {
	args := annotationArgs(ann, src)
	if args == nil {
		return 0, false
	}
	for i := 0; i < int(args.ChildCount()); i++ {
		child := args.Child(i)
		if child.Type() == "decimal_integer_literal" {
			return parseInt(nodeText(child, src)), true
		}
	}
	return 0, false
}

//ff:func feature=scan type=extract control=iteration dimension=1 topic=spring
//ff:what 어노테이션의 첫 번째 문자열 인자를 반환한다
package spring

import sitter "github.com/smacker/go-tree-sitter"

func firstStringArg(ann *sitter.Node, src []byte) string {
	args := annotationArgs(ann, src)
	if args == nil {
		return ""
	}
	for i := 0; i < int(args.ChildCount()); i++ {
		child := args.Child(i)
		if child.Type() == "string_literal" {
			return unquoteJava(nodeText(child, src))
		}
	}
	return ""
}

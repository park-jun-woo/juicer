//ff:func feature=scan type=extract control=iteration dimension=1 topic=dotnet
//ff:what 어트리뷰트 인자에서 문자열 리터럴을 찾는다
package dotnet

import sitter "github.com/smacker/go-tree-sitter"

func findStringLiteralInArg(arg *sitter.Node, src []byte) string {
	for i := 0; i < int(arg.ChildCount()); i++ {
		child := arg.Child(i)
		if child.Type() == "string_literal" {
			return unquoteCSharp(nodeText(child, src))
		}
	}
	return ""
}

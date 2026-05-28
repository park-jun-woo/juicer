//ff:func feature=scan type=extract control=iteration dimension=1 topic=dotnet
//ff:what argument_list에서 첫 번째 문자열 인자를 추출한다
package dotnet

import sitter "github.com/smacker/go-tree-sitter"

func extractFirstStringFromArgs(args *sitter.Node, src []byte) string {
	for _, arg := range childrenOfType(args, "argument") {
		strLit := findChildByType(arg, "string_literal")
		if strLit != nil {
			return unquoteCSharp(nodeText(strLit, src))
		}
	}
	return ""
}

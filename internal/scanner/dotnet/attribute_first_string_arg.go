//ff:func feature=scan type=extract control=iteration dimension=1 topic=dotnet
//ff:what 어트리뷰트의 첫 번째 문자열 인자를 반환한다
package dotnet

import sitter "github.com/smacker/go-tree-sitter"

func attributeFirstStringArg(attr *sitter.Node, src []byte) string {
	args := findChildByType(attr, "attribute_argument_list")
	if args == nil {
		return ""
	}
	for _, arg := range childrenOfType(args, "attribute_argument") {
		s := findStringLiteralInArg(arg, src)
		if s != "" {
			return s
		}
	}
	return ""
}

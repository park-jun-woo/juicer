//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what arguments에서 첫 인자의 문자열 내용을 반환한다(없으면 ok=false)
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func firstArgString(args *sitter.Node, src []byte) (string, bool) {
	argNodes := childrenOfType(args, "argument")
	if len(argNodes) == 0 {
		return "", false
	}
	return extractStringContent(argNodes[0], src), true
}

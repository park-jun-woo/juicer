//ff:func feature=scan type=extract control=iteration dimension=1 topic=dotnet
//ff:what 어트리뷰트 인자에서 정수 리터럴을 찾는다
package dotnet

import (
	"strconv"

	sitter "github.com/smacker/go-tree-sitter"
)

func findIntLiteralInArg(arg *sitter.Node, src []byte) (int, bool) {
	for i := 0; i < int(arg.ChildCount()); i++ {
		child := arg.Child(i)
		if child.Type() != "integer_literal" {
			continue
		}
		v, err := strconv.Atoi(nodeText(child, src))
		if err != nil {
			continue
		}
		return v, true
	}
	return 0, false
}

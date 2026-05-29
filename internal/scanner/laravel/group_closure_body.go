//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what ->group(...) 호출의 인자에서 클로저 본문(compound_statement)을 추출한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func groupClosureBody(mc *sitter.Node, fi fileInfo) *sitter.Node {
	groupArgs := findChildByType(mc, "arguments")
	if groupArgs == nil {
		return nil
	}
	return extractClosureBody(groupArgs, fi)
}

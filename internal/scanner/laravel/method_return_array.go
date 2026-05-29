//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what 메서드의 첫 return 문에서 첫 array_creation_expression 노드를 찾는다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func methodReturnArray(method *sitter.Node) *sitter.Node {
	retStmts := findAllByType(method, "return_statement")
	if len(retStmts) == 0 {
		return nil
	}
	arrNodes := findAllByType(retStmts[0], "array_creation_expression")
	if len(arrNodes) == 0 {
		return nil
	}
	return arrNodes[0]
}

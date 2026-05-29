//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what 클로저 인자에서 compound_statement(본문)를 추출한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// extractClosureBody extracts the compound_statement (body) from a closure argument.
func extractClosureBody(groupArgs *sitter.Node, fi fileInfo) *sitter.Node {
	closures := findAllByType(groupArgs, "anonymous_function_creation_expression")
	if len(closures) == 0 {
		closures = findAllByType(groupArgs, "arrow_function")
	}
	if len(closures) == 0 {
		return nil
	}
	return findChildByType(closures[0], "compound_statement")
}

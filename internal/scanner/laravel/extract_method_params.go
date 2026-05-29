//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what formal_parameters 노드에서 메서드 파라미터들을 추출한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// extractMethodParams extracts parameters from a formal_parameters node.
func extractMethodParams(formalParams *sitter.Node, src []byte) []methodParam {
	var params []methodParam
	for _, sp := range childrenOfType(formalParams, "simple_parameter") {
		params = append(params, parseSimpleParam(sp, src))
	}
	return params
}

//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 메서드 정의에서 공식 파라미터를 추출한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// extractMethodParams parses formal parameters of a method_definition.
func extractMethodParams(m *sitter.Node, src []byte) methodParams {
	var result methodParams
	params := findChildByType(m, "formal_parameters")
	if params == nil {
		return result
	}
	for i := 0; i < int(params.ChildCount()); i++ {
		child := params.Child(i)
		switch child.Type() {
		case "required_parameter", "optional_parameter":
			extractOneParam(child, src, &result)
		}
	}
	return result
}

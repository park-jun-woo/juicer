//ff:func feature=scan type=extract control=sequence topic=django
//ff:what 함수 정의를 @api_view 함수 뷰로 파싱한다
package django

import sitter "github.com/smacker/go-tree-sitter"

// parseFuncView parses a function definition as an @api_view function if it is one.
func parseFuncView(funcDef *sitter.Node, fi fileInfo) *funcViewInfo {
	nameNode := findChildByType(funcDef, "identifier")
	if nameNode == nil {
		return nil
	}
	methods := extractAPIViewDecorator(funcDef, fi.src)
	if methods == nil {
		return nil
	}
	return &funcViewInfo{
		name:    nodeText(nameNode, fi.src),
		methods: methods,
		file:    fi.relPath,
		line:    int(nameNode.StartPoint().Row) + 1,
	}
}

//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 함수 정의에서 @api_view 데코레이터를 추출한다
package django

import sitter "github.com/smacker/go-tree-sitter"

// extractAPIViewDecorator checks if a function has @api_view decorator and returns the methods list.
func extractAPIViewDecorator(funcDef *sitter.Node, src []byte) []string {
	parent := funcDef.Parent()
	if parent == nil || parent.Type() != "decorated_definition" {
		return nil
	}
	for _, dec := range childrenOfType(parent, "decorator") {
		methods := parseAPIViewDecoratorNode(dec, src)
		if methods != nil {
			return methods
		}
	}
	return nil
}

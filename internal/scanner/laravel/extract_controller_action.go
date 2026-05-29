//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what [Controller::class, 'method'] 배열에서 컨트롤러 클래스와 메서드를 추출한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// extractControllerAction extracts controller class and method from
// [Controller::class, 'method'] array syntax.
func extractControllerAction(node *sitter.Node, src []byte) (string, string) {
	arr := resolveArrayNode(node)
	if arr == nil {
		return "", ""
	}
	elems := childrenOfType(arr, "array_element_initializer")
	if len(elems) < 2 {
		return "", ""
	}
	controller := controllerClassName(elems[0], src)
	action := extractStringContent(elems[1], src)
	return controller, action
}

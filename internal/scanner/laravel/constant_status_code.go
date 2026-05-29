//ff:func feature=scan type=convert control=sequence topic=laravel
//ff:what Response::HTTP_* 상수 접근 노드에서 상태 코드를 매핑한다(아니면 빈 문자열)
package laravel

import sitter "github.com/smacker/go-tree-sitter"

func constantStatusCode(arg *sitter.Node, src []byte) string {
	cca := findChildByType(arg, "class_constant_access_expression")
	if cca == nil {
		return ""
	}
	names := childrenOfType(cca, "name")
	if len(names) == 0 {
		return ""
	}
	return httpStatusConstants[nodeText(names[len(names)-1], src)]
}

//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what 클래스의 extends 절에서 팩토리 함수 호출 이름을 추출한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// extendsFactoryName returns the function name from an extends clause that
// uses a factory-call pattern (e.g. `extends BaseController<...>(...)`).
// Returns "" if the class has no extends clause or the extends target is not
// a call_expression.
func extendsFactoryName(cls *sitter.Node, src []byte) string {
	heritage := findChildByType(cls, "class_heritage")
	if heritage == nil {
		return ""
	}
	ext := findChildByType(heritage, "extends_clause")
	if ext == nil {
		return ""
	}
	call := findChildByType(ext, "call_expression")
	if call == nil {
		return ""
	}
	ident := findChildByType(call, "identifier")
	if ident == nil {
		return ""
	}
	return nodeText(ident, src)
}

//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 자식 클래스의 extends 절에서 제네릭 타입 인자를 추출한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// extractTypeArgs extracts the type arguments from the extends clause of a
// child class. For example, given `extends BaseController<Category, CategoryDto>(...)`,
// it returns ["Category", "CategoryDto"].
//
// AST path: class_heritage -> extends_clause -> call_expression -> type_arguments.
func extractTypeArgs(cls *sitter.Node, src []byte) []string {
	heritage := findChildByType(cls, "class_heritage")
	if heritage == nil {
		return nil
	}
	ext := findChildByType(heritage, "extends_clause")
	if ext == nil {
		return nil
	}
	call := findChildByType(ext, "call_expression")
	if call == nil {
		return nil
	}
	typeArgs := findChildByType(call, "type_arguments")
	if typeArgs == nil {
		return nil
	}
	var result []string
	for i := 0; i < int(typeArgs.ChildCount()); i++ {
		child := typeArgs.Child(i)
		switch child.Type() {
		case "type_identifier", "predefined_type", "generic_type":
			result = append(result, nodeText(child, src))
		}
	}
	return result
}

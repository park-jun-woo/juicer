//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 내부 클래스의 type_parameters에서 제네릭 파라미터 이름을 추출한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// extractTypeParams extracts the type parameter names from a class declaration's
// type_parameters. For example, given `class Foo<D, B, CreateDtoType>`, it
// returns ["D", "B", "CreateDtoType"].
func extractTypeParams(cls *sitter.Node, src []byte) []string {
	typeParams := findChildByType(cls, "type_parameters")
	if typeParams == nil {
		return nil
	}
	var result []string
	for i := 0; i < int(typeParams.ChildCount()); i++ {
		child := typeParams.Child(i)
		if child.Type() != "type_parameter" {
			continue
		}
		// type_parameter's first type_identifier is the name
		ident := findChildByType(child, "type_identifier")
		if ident != nil {
			result = append(result, nodeText(ident, src))
		}
	}
	return result
}

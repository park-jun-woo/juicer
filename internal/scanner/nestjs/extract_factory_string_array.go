//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 팩토리 함수의 두 번째 인수에서 문자열 배열을 추출한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// extractFactoryStringArray extracts the second argument string array from factory calls.
// e.g. OmitType(Base, ['field1', 'field2']) returns ["field1", "field2"].
func extractFactoryStringArray(args *sitter.Node, src []byte) []string {
	var arr *sitter.Node
	for i := 0; i < int(args.ChildCount()); i++ {
		child := args.Child(i)
		// [values] as const → as_expression > array
		if child.Type() == "as_expression" {
			child = findChildByType(child, "array")
		}
		if child != nil && child.Type() == "array" {
			arr = child
			break
		}
	}
	if arr == nil {
		return nil
	}
	var result []string
	for i := 0; i < int(arr.ChildCount()); i++ {
		child := arr.Child(i)
		if child.Type() == "string" {
			result = append(result, unquoteTS(nodeText(child, src)))
		}
	}
	return result
}

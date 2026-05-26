//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 팩토리 함수 인수에서 베이스 클래스명을 추출한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// extractFactoryBaseClass extracts the base class name from factory arguments.
// e.g. arguments node of PartialType(CreateTaskDto) returns "CreateTaskDto".
func extractFactoryBaseClass(args *sitter.Node, src []byte) string {
	for i := 0; i < int(args.ChildCount()); i++ {
		child := args.Child(i)
		if child.Type() == "identifier" {
			return nodeText(child, src)
		}
	}
	return ""
}

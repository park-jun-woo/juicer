//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what argument_list에서 dependencies= 키워드의 리스트 값을 찾아 Depends 호출을 추출한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// findDependenciesKeyword finds the dependencies= keyword argument and
// extracts Depends() calls from its list value.
func findDependenciesKeyword(args *sitter.Node, src []byte) []string {
	for i := 0; i < int(args.ChildCount()); i++ {
		child := args.Child(i)
		if child.Type() != "keyword_argument" {
			continue
		}
		keyNode := findChildByType(child, "identifier")
		if keyNode == nil || nodeText(keyNode, src) != "dependencies" {
			continue
		}
		listNode := findChildByType(child, "list")
		if listNode == nil {
			continue
		}
		return extractDependsFromList(listNode, src)
	}
	return nil
}

//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what argument_list 에서 이름으로 키워드 인자 값을 추출한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// extractKeywordArg finds a keyword argument by name and returns its string value.
// e.g., from prefix="/users" returns "/users".
func extractKeywordArg(args *sitter.Node, name string, src []byte) string {
	for i := 0; i < int(args.ChildCount()); i++ {
		child := args.Child(i)
		if child.Type() != "keyword_argument" {
			continue
		}
		val := tryMatchKeyword(child, name, src)
		if val != "" {
			return val
		}
	}
	return ""
}

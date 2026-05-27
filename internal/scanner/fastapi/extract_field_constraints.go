//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what Field(...) 호출에서 ge, le, min_length, max_length 키워드 인자를 추출한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// extractFieldConstraints extracts ge, le, min_length, max_length keyword args from Field(...).
func extractFieldConstraints(args *sitter.Node, src []byte, f *pydanticField) {
	for i := 0; i < int(args.ChildCount()); i++ {
		child := args.Child(i)
		if child.Type() != "keyword_argument" {
			continue
		}
		keyNode := findChildByType(child, "identifier")
		if keyNode == nil {
			continue
		}
		key := nodeText(keyNode, src)
		valStr := valueAfterEquals(child, src)
		applyFieldConstraint(key, valStr, f)
	}
}

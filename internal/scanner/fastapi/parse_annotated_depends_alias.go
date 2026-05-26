//ff:func feature=scan type=parse control=sequence topic=fastapi
//ff:what assignment 노드가 Annotated[T, Depends(func)] 별칭인지 파싱한다
package fastapi

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// parseAnnotatedDependsAlias checks if an assignment node is of the form
// X = Annotated[T, Depends(func)] and returns (alias, dependsFuncName).
// Returns ("", "") if not matching.
func parseAnnotatedDependsAlias(assign *sitter.Node, src []byte) (string, string) {
	leftNode := findChildByType(assign, "identifier")
	if leftNode == nil {
		return "", ""
	}
	alias := nodeText(leftNode, src)

	sub := findChildByType(assign, "subscript")
	if sub == nil {
		return "", ""
	}
	subText := nodeText(sub, src)
	if !strings.HasPrefix(subText, "Annotated[") {
		return "", ""
	}

	return alias, extractDependsFromAnnotated(subText)
}

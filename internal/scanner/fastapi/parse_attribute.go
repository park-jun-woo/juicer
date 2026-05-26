//ff:func feature=scan type=parse control=sequence topic=fastapi
//ff:what attribute 노드에서 라우터 변수명과 HTTP 메서드를 파싱한다
package fastapi

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// parseAttribute parses "router.get" into ("router", "GET").
func parseAttribute(attrNode *sitter.Node, src []byte) (string, string) {
	attrText := nodeText(attrNode, src)
	parts := strings.SplitN(attrText, ".", 2)
	if len(parts) != 2 {
		return "", ""
	}
	httpMethod, ok := httpMethods[parts[1]]
	if !ok {
		return "", ""
	}
	return parts[0], httpMethod
}

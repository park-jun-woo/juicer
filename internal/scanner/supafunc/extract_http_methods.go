//ff:func feature=scan type=extract control=sequence topic=supafunc
//ff:what 리프 binary_expression에서 좌변이 req.method/method인 경우만 우변 문자열을 HTTP 메서드로 추출한다 (OPTIONS 제외)
package supafunc

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

func extractHTTPMethods(body *sitter.Node, src []byte) []string {
	var methods []string
	seen := map[string]bool{}
	walkNodes(body, func(n *sitter.Node) {
		if n.Type() != "binary_expression" {
			return
		}
		// 복합 binary_expression은 스킵 — 리프에서 처리한다
		if hasChildBinaryExpr(n) {
			return
		}
		left := n.ChildByFieldName("left")
		if left == nil {
			return
		}
		leftText := nodeText(left, src)
		if leftText != "req.method" && leftText != "method" {
			return
		}
		right := n.ChildByFieldName("right")
		if right == nil {
			return
		}
		strNodes := findAllByType(right, "string")
		for _, s := range strNodes {
			method := unquoteTS(nodeText(s, src))
			method = strings.ToUpper(method)
			if method == "OPTIONS" {
				continue
			}
			if !seen[method] {
				seen[method] = true
				methods = append(methods, method)
			}
		}
	})
	return methods
}

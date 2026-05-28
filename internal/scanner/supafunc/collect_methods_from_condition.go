//ff:func feature=scan type=extract control=sequence topic=supafunc
//ff:what 조건 노드에서 req.method === "X" 형태의 모든 리프 비교를 찾아 메서드 목록을 반환한다 (OR 조건 지원)
package supafunc

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

func collectMethodsFromCondition(cond *sitter.Node, src []byte) []string {
	seen := map[string]bool{}
	var methods []string
	walkNodes(cond, func(n *sitter.Node) {
		if n.Type() != "binary_expression" {
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
		op := operatorOfBinaryExpr(n, src)
		if op == "!==" || op == "!=" {
			return
		}
		right := n.ChildByFieldName("right")
		if right == nil {
			return
		}
		strNodes := findAllByType(right, "string")
		for _, s := range strNodes {
			method := strings.ToUpper(unquoteTS(nodeText(s, src)))
			if method == "OPTIONS" {
				continue
			}
			if method != "" && !seen[method] {
				seen[method] = true
				methods = append(methods, method)
			}
		}
	})
	return methods
}

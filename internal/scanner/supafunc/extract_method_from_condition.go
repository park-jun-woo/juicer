//ff:func feature=scan type=extract control=sequence topic=supafunc
//ff:what 조건식 노드에서 req.method 비교 대상 HTTP 메서드 문자열을 추출한다
package supafunc

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

func extractMethodFromCondition(cond *sitter.Node, src []byte) string {
	var method string
	walkNodes(cond, func(n *sitter.Node) {
		if method != "" {
			return
		}
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
		// 연산자가 부정(!=, !==)이면 가드 패턴 — 블록 등록하지 않음
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
			method = strings.ToUpper(unquoteTS(nodeText(s, src)))
			return
		}
	})
	return method
}

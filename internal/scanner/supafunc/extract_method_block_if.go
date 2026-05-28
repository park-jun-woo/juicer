//ff:func feature=scan type=extract control=iteration dimension=1 topic=supafunc
//ff:what if(req.method==="X") 분기에서 메서드명과 consequence 블록을 매핑한다
package supafunc

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

func extractMethodBlockFromIf(ifNode *sitter.Node, src []byte, result map[string]*sitter.Node) {
	cond := ifNode.ChildByFieldName("condition")
	if cond == nil {
		return
	}
	condText := nodeText(cond, src)
	if !strings.Contains(condText, "req.method") && !strings.Contains(condText, "method") {
		return
	}
	consequence := ifNode.ChildByFieldName("consequence")
	if consequence == nil {
		return
	}
	// OR 조건에서 다중 메서드 추출: 모든 리프 binary_expression을 순회
	methods := collectMethodsFromCondition(cond, src)
	for _, m := range methods {
		// 이미 등록된 메서드는 덮어쓰지 않는다.
		// OR 조건 블록(body 파싱)이 먼저 등록된 경우 단일 메서드 블록이 덮지 않도록 한다.
		if _, exists := result[m]; !exists {
			result[m] = consequence
		}
	}
}

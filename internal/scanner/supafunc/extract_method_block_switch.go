//ff:func feature=scan type=extract control=iteration dimension=1 topic=supafunc
//ff:what switch(req.method) 분기에서 각 case의 메서드명과 블록 노드를 매핑한다
package supafunc

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

func extractMethodBlockFromSwitch(switchNode *sitter.Node, src []byte, result map[string]*sitter.Node) {
	value := switchNode.ChildByFieldName("value")
	if value == nil {
		return
	}
	valueText := nodeText(value, src)
	if !strings.Contains(valueText, "req.method") && !strings.Contains(valueText, "method") {
		return
	}
	body := switchNode.ChildByFieldName("body")
	if body == nil {
		return
	}
	cases := childrenOfType(body, "switch_case")
	for _, sc := range cases {
		caseValue := sc.ChildByFieldName("value")
		if caseValue == nil {
			continue
		}
		method := strings.ToUpper(unquoteTS(nodeText(caseValue, src)))
		if method == "OPTIONS" {
			continue
		}
		result[method] = sc
	}
}

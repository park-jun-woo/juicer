//ff:func feature=scan type=extract control=sequence topic=supafunc
//ff:what 콜백 본문의 모든 req.json() 호출이 메서드 블록 바이트 범위 안에 있는지 확인한다
package supafunc

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

func allJSONInsideBlocks(callbackBody *sitter.Node, src []byte, methodBlocks map[string]*sitter.Node) bool {
	found := false
	allInside := true
	walkNodes(callbackBody, func(n *sitter.Node) {
		if !allInside {
			return
		}
		// req.json()이 포함된 선언문을 찾는다
		if n.Type() != "lexical_declaration" && n.Type() != "variable_declaration" {
			return
		}
		text := nodeText(n, src)
		if !strings.Contains(text, "req.json()") {
			return
		}
		found = true
		start := n.StartByte()
		end := n.EndByte()
		inside := false
		for _, block := range methodBlocks {
			if start >= block.StartByte() && end <= block.EndByte() {
				inside = true
				break
			}
		}
		if !inside {
			allInside = false
		}
	})
	// req.json()이 하나도 없으면 블록 밖에 body가 없다는 의미
	if !found {
		return true
	}
	return allInside
}

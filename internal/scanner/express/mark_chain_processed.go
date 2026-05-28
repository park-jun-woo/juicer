//ff:func feature=scan type=extract control=sequence topic=express
//ff:what 체인 내 모든 call_expression을 처리 완료로 표시한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func markChainProcessed(node *sitter.Node, processed map[uintptr]bool) {
	walkNodes(node, func(n *sitter.Node) {
		if n.Type() == "call_expression" {
			processed[uintptr(n.StartByte())] = true
		}
	})
}

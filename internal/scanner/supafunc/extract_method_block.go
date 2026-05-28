//ff:func feature=scan type=extract control=sequence topic=supafunc
//ff:what if(req.method==="X") 또는 switch(req.method) 분기에서 메서드별 블록 노드를 추출한다
package supafunc

import sitter "github.com/smacker/go-tree-sitter"

func extractMethodBlock(callbackBody *sitter.Node, src []byte) map[string]*sitter.Node {
	result := map[string]*sitter.Node{}

	walkNodes(callbackBody, func(n *sitter.Node) {
		switch n.Type() {
		case "if_statement":
			extractMethodBlockFromIf(n, src, result)
		case "switch_statement":
			extractMethodBlockFromSwitch(n, src, result)
		}
	})

	return result
}

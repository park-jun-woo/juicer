//ff:func feature=scan type=extract control=sequence topic=express
//ff:what res.sendStatus(N) 호출에서 인자 N을 문자열로 추출한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func extractSendStatusArg(call *sitter.Node, src []byte) string {
	args := findChildByType(call, "arguments")
	if args == nil {
		return ""
	}
	argNodes := collectArgNodes(args)
	if len(argNodes) == 0 {
		return ""
	}
	first := argNodes[0]
	if first.Type() == "number" {
		return nodeText(first, src)
	}
	return ""
}

//ff:func feature=scan type=extract control=iteration dimension=1 topic=supafunc
//ff:what searchParams.get() 호출 인자에서 문자열 파라미터 이름을 추출한다
package supafunc

import sitter "github.com/smacker/go-tree-sitter"

func extractParamFromCall(call *sitter.Node, src []byte) string {
	args := findChildByType(call, "arguments")
	if args == nil {
		return ""
	}
	strNodes := childrenOfType(args, "string")
	for _, s := range strNodes {
		name := unquoteTS(nodeText(s, src))
		if name != "" {
			return name
		}
	}
	return ""
}

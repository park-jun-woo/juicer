//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what abort 계열 호출에서 상태 코드 인자를 해석해 반환한다(아니면 빈 문자열)
package laravel

import sitter "github.com/smacker/go-tree-sitter"

func abortCallStatus(call *sitter.Node, src []byte) string {
	nameNode := findChildByType(call, "name")
	if nameNode == nil {
		return ""
	}
	idx, ok := abortFunctions[nodeText(nameNode, src)]
	if !ok {
		return ""
	}
	args := findChildByType(call, "arguments")
	if args == nil {
		return ""
	}
	argList := childrenOfType(args, "argument")
	if idx >= len(argList) {
		return ""
	}
	code := resolveStatusArg(argList[idx], src)
	if !isNumericStatus(code) {
		return ""
	}
	return code
}

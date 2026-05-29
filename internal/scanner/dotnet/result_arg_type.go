//ff:func feature=scan type=extract control=sequence topic=dotnet
//ff:what 팩토리 호출의 첫 인자에서 응답 본문 타입을 추출한다
package dotnet

import sitter "github.com/smacker/go-tree-sitter"

func resultArgType(args *sitter.Node, src []byte) (string, bool) {
	if args == nil {
		return "", false
	}
	argNodes := childrenOfType(args, "argument")
	if len(argNodes) == 0 {
		return "", false
	}
	return argType(argNodes[0], src)
}

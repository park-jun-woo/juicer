//ff:func feature=scan type=extract control=sequence topic=dotnet
//ff:what StatusCode(<int>, body) 호출에서 정수 상태 코드와 본문 타입을 추출한다
package dotnet

import (
	"strconv"

	sitter "github.com/smacker/go-tree-sitter"
)

func matchStatusCodeInvocation(args *sitter.Node, src []byte) bodyResponse {
	if args == nil {
		return bodyResponse{}
	}
	argNodes := childrenOfType(args, "argument")
	if len(argNodes) == 0 {
		return bodyResponse{}
	}
	lit := findChildByType(argNodes[0], "integer_literal")
	if lit == nil {
		return bodyResponse{}
	}
	if _, err := strconv.Atoi(nodeText(lit, src)); err != nil {
		return bodyResponse{}
	}
	res := bodyResponse{status: nodeText(lit, src), found: true}
	if len(argNodes) > 1 {
		res.typeName, res.isArray = argType(argNodes[1], src)
	}
	return res
}

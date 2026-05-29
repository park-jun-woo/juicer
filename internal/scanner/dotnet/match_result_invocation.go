//ff:func feature=scan type=extract control=sequence topic=dotnet
//ff:what IActionResult 팩토리 호출(Ok/Created/StatusCode 등)에서 상태·타입을 추출한다
package dotnet

import sitter "github.com/smacker/go-tree-sitter"

func matchResultInvocation(inv *sitter.Node, src []byte) bodyResponse {
	name := findChildByType(inv, "identifier")
	if name == nil {
		return bodyResponse{}
	}
	args := findChildByType(inv, "argument_list")
	method := nodeText(name, src)
	if method == "StatusCode" {
		return matchStatusCodeInvocation(args, src)
	}
	status, ok := resultsStatusMethods[method]
	if !ok {
		return bodyResponse{}
	}
	res := bodyResponse{status: status, found: true}
	res.typeName, res.isArray = resultArgType(args, src)
	return res
}

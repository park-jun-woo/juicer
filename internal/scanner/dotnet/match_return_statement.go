//ff:func feature=scan type=extract control=sequence topic=dotnet
//ff:what return_statement 안의 invocation_expression을 응답 정보로 매칭한다
package dotnet

import sitter "github.com/smacker/go-tree-sitter"

func matchReturnStatement(ret *sitter.Node, src []byte) bodyResponse {
	inv := findChildByType(ret, "invocation_expression")
	if inv == nil {
		return bodyResponse{}
	}
	return matchResultInvocation(inv, src)
}

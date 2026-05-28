//ff:func feature=scan type=extract control=sequence topic=express
//ff:what res.status(N).json() 체인에서 상태 코드 N을 추출한다
package express

import sitter "github.com/smacker/go-tree-sitter"

// extractStatusFromChain — call_expression(예: res.status(201).json(...))에서
// 내부 res.status(N) 호출의 인자 N을 문자열로 반환한다.
// 체인이 아니면 빈 문자열을 반환한다.
func extractStatusFromChain(call *sitter.Node, src []byte) string {
	mem := findChildByType(call, "member_expression")
	if mem == nil {
		return ""
	}
	obj := mem.ChildByFieldName("object")
	if obj == nil || obj.Type() != "call_expression" {
		return ""
	}
	if !isResStatusCall(obj, src) {
		return ""
	}
	// obj = res.status(N) — arguments에서 첫 번째 number를 추출
	args := findChildByType(obj, "arguments")
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

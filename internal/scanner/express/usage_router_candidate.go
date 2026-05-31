//ff:func feature=scan type=extract control=sequence topic=express
//ff:what call_expression이 `<ident>.<httpMethod>("/path", …)` 폴백 라우터 패턴이면 라우터 변수명을 반환한다 (아니면 "")
package express

import sitter "github.com/smacker/go-tree-sitter"

var usageHTTPMethods = map[string]bool{
	"get": true, "post": true, "put": true, "delete": true,
	"patch": true, "options": true, "head": true, "all": true,
	"use": true,
}

func usageRouterCandidate(call *sitter.Node, src []byte) string {
	mem := findChildByType(call, "member_expression")
	if mem == nil {
		return ""
	}
	prop := mem.ChildByFieldName("property")
	if prop == nil || !usageHTTPMethods[nodeText(prop, src)] {
		return ""
	}
	obj := findChildByType(mem, "identifier")
	if obj == nil {
		return ""
	}
	varName := nodeText(obj, src)
	if varName == "express" || varName == "module" || varName == "exports" {
		return ""
	}
	args := findChildByType(call, "arguments")
	if args == nil || args.NamedChildCount() == 0 {
		return ""
	}
	firstArg := args.NamedChild(0)
	if firstArg == nil || firstArg.Type() != "string" {
		return ""
	}
	// path 형태 가드: 첫 문자열 인자가 라우트 path 형태(`/` 시작 또는 `*` catch-all)일 때만
	// 라우터 후보로 인정한다. 이로써 `req.get('user-agent')`·`config.get('urls')`·
	// `model.get('status')` 같은 비-HTTP 메서드 호출이 usage 폴백을 통과하지 못한다
	// (블랙리스트로 막을 수 없던 임의 변수명 해소).
	if !isRoutePathArg(unquoteTS(nodeText(firstArg, src))) {
		return ""
	}
	return varName
}

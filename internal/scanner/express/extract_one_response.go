//ff:func feature=scan type=extract control=selection topic=express
//ff:what 단일 call_expression에서 res.json()/res.send()/res.sendStatus()/res.status().json() 응답을 추출한다
package express

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
)

// extractOneResponse — call_expression을 받아 Express 응답 패턴을 분석한다.
//
// 패턴:
//
//	res.json(...)               → 200 / json
//	res.status(N).json(...)     → N   / json
//	res.sendStatus(N)           → N   / empty
//	res.send(...)               → 200 / text
//	res.status(N).send(...)     → N   / text
//	res.render(...)             → 200 / html
//	res.status(N).render(...)   → N   / html
//	res.redirect(...)           → 302 / empty
//	res.status(N).redirect(...) → N   / empty
func extractOneResponse(call *sitter.Node, src []byte) *scanner.Response {
	methodName, ok := isResMethodCall(call, src)
	if !ok {
		return nil
	}

	switch methodName {
	case "json":
		status := extractStatusFromChain(call, src)
		if status == "" {
			status = "200"
		}
		return &scanner.Response{Status: status, Kind: "json"}

	case "send":
		status := extractStatusFromChain(call, src)
		if status == "" {
			status = "200"
		}
		return &scanner.Response{Status: status, Kind: "text"}

	case "sendStatus":
		status := extractSendStatusArg(call, src)
		if status == "" {
			status = "200"
		}
		return &scanner.Response{Status: status, Kind: "empty"}

	case "render":
		status := extractStatusFromChain(call, src)
		if status == "" {
			status = "200"
		}
		return &scanner.Response{Status: status, Kind: "html"}

	case "redirect":
		// res.redirect는 기본 302. res.status(N).redirect(...)면 체인 상태코드 반영.
		status := extractStatusFromChain(call, src)
		if status == "" {
			status = "302"
		}
		return &scanner.Response{Status: status, Kind: "empty"}
	}

	return nil
}

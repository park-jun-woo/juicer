//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what HttpResponse::<status> 노드 하나를 해석해 응답 목록에 추가한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func captureResponse(n *sitter.Node, ctx *responseCtx) {
	if n.Type() != "scoped_identifier" {
		return
	}
	parts := splitScoped(nodeText(n, ctx.src))
	if len(parts) != 2 || parts[0] != "HttpResponse" {
		return
	}
	code, ok := httpResponseStatuses[parts[1]]
	if !ok || ctx.seen[code] {
		return
	}
	ctx.seen[code] = true
	resp := scanner.Response{Status: code, Kind: detectResponseKind(n, ctx.src)}
	if resp.Kind == "json" {
		applyResponseBody(&resp, n, ctx)
	}
	ctx.responses = append(ctx.responses, resp)
}
